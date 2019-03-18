package s3

import (
	"net/http"
	"regexp"
)

// note: the repo matcher is stricter than what PFS enforces, because this is
// based off of minio's bucket matching regexp. At the point where we're
// matching errors, we shouldn't be attempting to operate on buckets that
// don't match the minio regexp anyways.
var (
	repoNotFoundMatcher   = regexp.MustCompile(`repos/ ?[a-zA-Z0-9.\-_]{1,255} not found`)
	branchNotFoundMatcher = regexp.MustCompile(`branches/[a-zA-Z0-9.\-_]{1,255}/ [^ ]+ not found`)
	fileNotFoundMatcher   = regexp.MustCompile(`file .+ not found`)
)

// Error is an XML-encodable error response
type Error struct {
	httpStatus int    `xml"-"`
	Code       string `xml:"Code"`
	Message    string `xml:"Message"`
	Resource   string `xml:"Resource"`
	RequestID  string `xml:"RequestId"`
}

func writeError(w http.ResponseWriter, r *http.Request, httpStatus int, code string, message string) {
	writeXML(w, r, httpStatus, &Error{
		httpStatus: httpStatus,
		Code:       code,
		Message:    message,
		Resource:   r.URL.Path,
		RequestID:  r.Header.Get("X-Request-ID"),
	})
}

func badDigestError(w http.ResponseWriter, r *http.Request) {
	writeError(w, r, http.StatusBadRequest, "BadDigest", "The Content-MD5 you specified did not match what we received.")
}

func bucketNotEmptyError(w http.ResponseWriter, r *http.Request) {
	writeError(w, r, http.StatusConflict, "BucketNotEmpty", "The bucket you tried to delete is not empty.")
}

func bucketAlreadyExistsError(w http.ResponseWriter, r *http.Request) {
	writeError(w, r, http.StatusBadRequest, "BucketAlreadyExists", "There is already a repo with that name.")
}

func internalError(w http.ResponseWriter, r *http.Request, err error) {
	writeError(w, r, http.StatusInternalServerError, "InternalError", err.Error())
}

func invalidBucketNameError(w http.ResponseWriter, r *http.Request) {
	writeError(w, r, http.StatusBadRequest, "InvalidBucketName", "The specified repo or branch either has an invalid name, or is not serviceable.")
}

func invalidArgument(w http.ResponseWriter, r *http.Request) {
	writeError(w, r, http.StatusBadRequest, "InvalidArgument", "Invalid Argument")
}

// note: this is not a standard s3 error
func invalidDelimiterError(w http.ResponseWriter, r *http.Request) {
	writeError(w, r, http.StatusBadRequest, "InvalidDelimiter", "The delimiter you specified is invalid. It must be '' or '/'.")
}

func invalidDigestError(w http.ResponseWriter, r *http.Request) {
	writeError(w, r, http.StatusBadRequest, "InvalidDigest", "The Content-MD5 you specified is not valid.")
}

// note: this is not a standard s3 error
func invalidFilePathError(w http.ResponseWriter, r *http.Request) {
	writeError(w, r, http.StatusBadRequest, "InvalidFilePath", "Invalid file path")
}

func malformedXMLError(w http.ResponseWriter, r *http.Request) {
	writeError(w, r, http.StatusBadRequest, "MalformedXML", "The XML you provided was not well-formed or would not validate against S3's published schema.")
}

func methodNotAllowedError(w http.ResponseWriter, r *http.Request) {
	writeError(w, r, http.StatusMethodNotAllowed, "MethodNotAllowed", "The specified method is not allowed against this resource.")
}

func noSuchBucketError(w http.ResponseWriter, r *http.Request) {
	writeError(w, r, http.StatusNotFound, "NoSuchBucket", "The specified bucket does not exist.")
}

func noSuchKeyError(w http.ResponseWriter, r *http.Request) {
	writeError(w, r, http.StatusNotFound, "NoSuchKey", "The specified key does not exist.")
}

func notFoundError(w http.ResponseWriter, r *http.Request, err error) {
	s := err.Error()

	if repoNotFoundMatcher.MatchString(s) || branchNotFoundMatcher.MatchString(s) {
		noSuchBucketError(w, r)
	} else if fileNotFoundMatcher.MatchString(s) {
		noSuchKeyError(w, r)
	} else {
		internalError(w, r, err)
	}
}

func notImplementedError(w http.ResponseWriter, r *http.Request) {
	writeError(w, r, http.StatusNotImplemented, "NotImplemented", "This functionality is not implemented in the pachyderm s3gateway.")
}
