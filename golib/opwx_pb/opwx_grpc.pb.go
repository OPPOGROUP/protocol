// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: proto/opwx.proto

package opwx_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Opwx_PushCheckinResults_FullMethodName = "/proto.Opwx/PushCheckinResults"
)

// OpwxClient is the client API for Opwx service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpwxClient interface {
	PushCheckinResults(ctx context.Context, in *CheckinResults, opts ...grpc.CallOption) (*PushResponse, error)
}

type opwxClient struct {
	cc grpc.ClientConnInterface
}

func NewOpwxClient(cc grpc.ClientConnInterface) OpwxClient {
	return &opwxClient{cc}
}

func (c *opwxClient) PushCheckinResults(ctx context.Context, in *CheckinResults, opts ...grpc.CallOption) (*PushResponse, error) {
	out := new(PushResponse)
	err := c.cc.Invoke(ctx, Opwx_PushCheckinResults_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpwxServer is the server API for Opwx service.
// All implementations must embed UnimplementedOpwxServer
// for forward compatibility
type OpwxServer interface {
	PushCheckinResults(context.Context, *CheckinResults) (*PushResponse, error)
	mustEmbedUnimplementedOpwxServer()
}

// UnimplementedOpwxServer must be embedded to have forward compatible implementations.
type UnimplementedOpwxServer struct {
}

func (UnimplementedOpwxServer) PushCheckinResults(context.Context, *CheckinResults) (*PushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushCheckinResults not implemented")
}
func (UnimplementedOpwxServer) mustEmbedUnimplementedOpwxServer() {}

// UnsafeOpwxServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpwxServer will
// result in compilation errors.
type UnsafeOpwxServer interface {
	mustEmbedUnimplementedOpwxServer()
}

func RegisterOpwxServer(s grpc.ServiceRegistrar, srv OpwxServer) {
	s.RegisterService(&Opwx_ServiceDesc, srv)
}

func _Opwx_PushCheckinResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckinResults)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpwxServer).PushCheckinResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Opwx_PushCheckinResults_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpwxServer).PushCheckinResults(ctx, req.(*CheckinResults))
	}
	return interceptor(ctx, in, info, handler)
}

// Opwx_ServiceDesc is the grpc.ServiceDesc for Opwx service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Opwx_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Opwx",
	HandlerType: (*OpwxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushCheckinResults",
			Handler:    _Opwx_PushCheckinResults_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/opwx.proto",
}