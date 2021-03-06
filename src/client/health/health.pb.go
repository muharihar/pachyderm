// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: client/health/health.proto

package health

import (
	context "context"
	fmt "fmt"
	types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("client/health/health.proto", fileDescriptor_6a4830d93d85d5ea) }

var fileDescriptor_6a4830d93d85d5ea = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0xd1, 0xcf, 0x48, 0x4d, 0xcc, 0x29, 0xc9, 0x80, 0x52, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x94, 0x74, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x58, 0x34,
	0xa9, 0x34, 0x4d, 0x3f, 0x35, 0xb7, 0xa0, 0xa4, 0x12, 0xa2, 0xc8, 0xc8, 0x85, 0x8b, 0xcd, 0x03,
	0xac, 0x4c, 0xc8, 0x0a, 0xce, 0x12, 0xd3, 0x83, 0xe8, 0xd0, 0x83, 0xe9, 0xd0, 0x73, 0x05, 0xe9,
	0x90, 0xc2, 0x21, 0xae, 0xc4, 0xe0, 0xe4, 0x74, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c,
	0x0f, 0x1e, 0xc9, 0x31, 0xce, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x90, 0x9e, 0x59, 0x92, 0x51, 0x9a,
	0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x5f, 0x90, 0x98, 0x9c, 0x51, 0x99, 0x92, 0x5a, 0x84, 0xcc, 0x2a,
	0x2e, 0x4a, 0xd6, 0x47, 0x71, 0x7b, 0x12, 0x1b, 0xd8, 0x54, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x54, 0x45, 0xdf, 0x02, 0xd3, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HealthClient is the client API for Health service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthClient interface {
	Health(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Health(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/health.Health/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServer is the server API for Health service.
type HealthServer interface {
	Health(context.Context, *types.Empty) (*types.Empty, error)
}

// UnimplementedHealthServer can be embedded to have forward compatible implementations.
type UnimplementedHealthServer struct {
}

func (*UnimplementedHealthServer) Health(ctx context.Context, req *types.Empty) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health.Health/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Health(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "health.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _Health_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client/health/health.proto",
}
