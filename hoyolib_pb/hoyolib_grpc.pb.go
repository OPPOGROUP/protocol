// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: hoyolib_pb/hoyolib.proto

package hoyolib_pb

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
	Hoyolib_Register_FullMethodName       = "/Hoyolib/Register"
	Hoyolib_CheckIn_FullMethodName        = "/Hoyolib/CheckIn"
	Hoyolib_GetAccountInfo_FullMethodName = "/Hoyolib/GetAccountInfo"
)

// HoyolibClient is the client API for Hoyolib service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HoyolibClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	CheckIn(ctx context.Context, in *CheckInRequest, opts ...grpc.CallOption) (*CheckInResponse, error)
	GetAccountInfo(ctx context.Context, in *AccountInfoRequest, opts ...grpc.CallOption) (*AccountInfoResponse, error)
}

type hoyolibClient struct {
	cc grpc.ClientConnInterface
}

func NewHoyolibClient(cc grpc.ClientConnInterface) HoyolibClient {
	return &hoyolibClient{cc}
}

func (c *hoyolibClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, Hoyolib_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hoyolibClient) CheckIn(ctx context.Context, in *CheckInRequest, opts ...grpc.CallOption) (*CheckInResponse, error) {
	out := new(CheckInResponse)
	err := c.cc.Invoke(ctx, Hoyolib_CheckIn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hoyolibClient) GetAccountInfo(ctx context.Context, in *AccountInfoRequest, opts ...grpc.CallOption) (*AccountInfoResponse, error) {
	out := new(AccountInfoResponse)
	err := c.cc.Invoke(ctx, Hoyolib_GetAccountInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HoyolibServer is the server API for Hoyolib service.
// All implementations must embed UnimplementedHoyolibServer
// for forward compatibility
type HoyolibServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	CheckIn(context.Context, *CheckInRequest) (*CheckInResponse, error)
	GetAccountInfo(context.Context, *AccountInfoRequest) (*AccountInfoResponse, error)
	mustEmbedUnimplementedHoyolibServer()
}

// UnimplementedHoyolibServer must be embedded to have forward compatible implementations.
type UnimplementedHoyolibServer struct {
}

func (UnimplementedHoyolibServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedHoyolibServer) CheckIn(context.Context, *CheckInRequest) (*CheckInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIn not implemented")
}
func (UnimplementedHoyolibServer) GetAccountInfo(context.Context, *AccountInfoRequest) (*AccountInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountInfo not implemented")
}
func (UnimplementedHoyolibServer) mustEmbedUnimplementedHoyolibServer() {}

// UnsafeHoyolibServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HoyolibServer will
// result in compilation errors.
type UnsafeHoyolibServer interface {
	mustEmbedUnimplementedHoyolibServer()
}

func RegisterHoyolibServer(s grpc.ServiceRegistrar, srv HoyolibServer) {
	s.RegisterService(&Hoyolib_ServiceDesc, srv)
}

func _Hoyolib_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoyolibServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hoyolib_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoyolibServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hoyolib_CheckIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoyolibServer).CheckIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hoyolib_CheckIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoyolibServer).CheckIn(ctx, req.(*CheckInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hoyolib_GetAccountInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoyolibServer).GetAccountInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hoyolib_GetAccountInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoyolibServer).GetAccountInfo(ctx, req.(*AccountInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Hoyolib_ServiceDesc is the grpc.ServiceDesc for Hoyolib service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hoyolib_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Hoyolib",
	HandlerType: (*HoyolibServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Hoyolib_Register_Handler,
		},
		{
			MethodName: "CheckIn",
			Handler:    _Hoyolib_CheckIn_Handler,
		},
		{
			MethodName: "GetAccountInfo",
			Handler:    _Hoyolib_GetAccountInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hoyolib_pb/hoyolib.proto",
}

const (
	Opwx_PushCheckinResults_FullMethodName = "/Opwx/PushCheckinResults"
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
	ServiceName: "Opwx",
	HandlerType: (*OpwxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushCheckinResults",
			Handler:    _Opwx_PushCheckinResults_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hoyolib_pb/hoyolib.proto",
}