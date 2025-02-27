// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.9
// source: proto/user/user.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserService_CreateLimit_FullMethodName                   = "/user.UserService/CreateLimit"
	UserService_UpdateLimit_FullMethodName                   = "/user.UserService/UpdateLimit"
	UserService_UpdateParentLimit_FullMethodName             = "/user.UserService/UpdateParentLimit"
	UserService_HandleExpiredUserSubscription_FullMethodName = "/user.UserService/HandleExpiredUserSubscription"
	UserService_CheckAuthApiV1_FullMethodName                = "/user.UserService/CheckAuthApiV1"
	UserService_CheckAuthApiV2_FullMethodName                = "/user.UserService/CheckAuthApiV2"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateLimit(ctx context.Context, in *LimitRequest, opts ...grpc.CallOption) (*LimitResponse, error)
	UpdateLimit(ctx context.Context, in *LimitRequest, opts ...grpc.CallOption) (*Response, error)
	UpdateParentLimit(ctx context.Context, in *LimitRequest, opts ...grpc.CallOption) (*Response, error)
	HandleExpiredUserSubscription(ctx context.Context, in *ExpiredUserRequest, opts ...grpc.CallOption) (*Response, error)
	CheckAuthApiV1(ctx context.Context, in *CheckAuthRequest, opts ...grpc.CallOption) (*CheckAuthV1Response, error)
	CheckAuthApiV2(ctx context.Context, in *CheckAuthRequest, opts ...grpc.CallOption) (*CheckAuthV2Response, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateLimit(ctx context.Context, in *LimitRequest, opts ...grpc.CallOption) (*LimitResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LimitResponse)
	err := c.cc.Invoke(ctx, UserService_CreateLimit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateLimit(ctx context.Context, in *LimitRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, UserService_UpdateLimit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateParentLimit(ctx context.Context, in *LimitRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, UserService_UpdateParentLimit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) HandleExpiredUserSubscription(ctx context.Context, in *ExpiredUserRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, UserService_HandleExpiredUserSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CheckAuthApiV1(ctx context.Context, in *CheckAuthRequest, opts ...grpc.CallOption) (*CheckAuthV1Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckAuthV1Response)
	err := c.cc.Invoke(ctx, UserService_CheckAuthApiV1_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CheckAuthApiV2(ctx context.Context, in *CheckAuthRequest, opts ...grpc.CallOption) (*CheckAuthV2Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckAuthV2Response)
	err := c.cc.Invoke(ctx, UserService_CheckAuthApiV2_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	CreateLimit(context.Context, *LimitRequest) (*LimitResponse, error)
	UpdateLimit(context.Context, *LimitRequest) (*Response, error)
	UpdateParentLimit(context.Context, *LimitRequest) (*Response, error)
	HandleExpiredUserSubscription(context.Context, *ExpiredUserRequest) (*Response, error)
	CheckAuthApiV1(context.Context, *CheckAuthRequest) (*CheckAuthV1Response, error)
	CheckAuthApiV2(context.Context, *CheckAuthRequest) (*CheckAuthV2Response, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) CreateLimit(context.Context, *LimitRequest) (*LimitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLimit not implemented")
}
func (UnimplementedUserServiceServer) UpdateLimit(context.Context, *LimitRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLimit not implemented")
}
func (UnimplementedUserServiceServer) UpdateParentLimit(context.Context, *LimitRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParentLimit not implemented")
}
func (UnimplementedUserServiceServer) HandleExpiredUserSubscription(context.Context, *ExpiredUserRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleExpiredUserSubscription not implemented")
}
func (UnimplementedUserServiceServer) CheckAuthApiV1(context.Context, *CheckAuthRequest) (*CheckAuthV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAuthApiV1 not implemented")
}
func (UnimplementedUserServiceServer) CheckAuthApiV2(context.Context, *CheckAuthRequest) (*CheckAuthV2Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAuthApiV2 not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateLimit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateLimit(ctx, req.(*LimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateLimit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateLimit(ctx, req.(*LimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateParentLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateParentLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateParentLimit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateParentLimit(ctx, req.(*LimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_HandleExpiredUserSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpiredUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).HandleExpiredUserSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_HandleExpiredUserSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).HandleExpiredUserSubscription(ctx, req.(*ExpiredUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CheckAuthApiV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CheckAuthApiV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CheckAuthApiV1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CheckAuthApiV1(ctx, req.(*CheckAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CheckAuthApiV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CheckAuthApiV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CheckAuthApiV2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CheckAuthApiV2(ctx, req.(*CheckAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLimit",
			Handler:    _UserService_CreateLimit_Handler,
		},
		{
			MethodName: "UpdateLimit",
			Handler:    _UserService_UpdateLimit_Handler,
		},
		{
			MethodName: "UpdateParentLimit",
			Handler:    _UserService_UpdateParentLimit_Handler,
		},
		{
			MethodName: "HandleExpiredUserSubscription",
			Handler:    _UserService_HandleExpiredUserSubscription_Handler,
		},
		{
			MethodName: "CheckAuthApiV1",
			Handler:    _UserService_CheckAuthApiV1_Handler,
		},
		{
			MethodName: "CheckAuthApiV2",
			Handler:    _UserService_CheckAuthApiV2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user/user.proto",
}
