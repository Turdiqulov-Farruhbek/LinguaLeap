// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: user_progress.proto

package learn

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	UserProgressService_AddUserProgress_FullMethodName = "/languagelearning.UserProgressService/AddUserProgress"
	UserProgressService_GetUserProgress_FullMethodName = "/languagelearning.UserProgressService/GetUserProgress"
)

// UserProgressServiceClient is the client API for UserProgressService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// UserProgress xizmatini aniqlash
type UserProgressServiceClient interface {
	// Foydalanuvchi taraqqiyotini qo'shish uchun RPC
	AddUserProgress(ctx context.Context, in *AddUserProgressRequest, opts ...grpc.CallOption) (*AddUserProgressResponse, error)
	// Foydalanuvchi taraqqiyotini olish uchun RPC
	GetUserProgress(ctx context.Context, in *GetUserProgressRequest, opts ...grpc.CallOption) (*GetUserProgressResponse, error)
}

type userProgressServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserProgressServiceClient(cc grpc.ClientConnInterface) UserProgressServiceClient {
	return &userProgressServiceClient{cc}
}

func (c *userProgressServiceClient) AddUserProgress(ctx context.Context, in *AddUserProgressRequest, opts ...grpc.CallOption) (*AddUserProgressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddUserProgressResponse)
	err := c.cc.Invoke(ctx, UserProgressService_AddUserProgress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProgressServiceClient) GetUserProgress(ctx context.Context, in *GetUserProgressRequest, opts ...grpc.CallOption) (*GetUserProgressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserProgressResponse)
	err := c.cc.Invoke(ctx, UserProgressService_GetUserProgress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserProgressServiceServer is the server API for UserProgressService service.
// All implementations must embed UnimplementedUserProgressServiceServer
// for forward compatibility
//
// UserProgress xizmatini aniqlash
type UserProgressServiceServer interface {
	// Foydalanuvchi taraqqiyotini qo'shish uchun RPC
	AddUserProgress(context.Context, *AddUserProgressRequest) (*AddUserProgressResponse, error)
	// Foydalanuvchi taraqqiyotini olish uchun RPC
	GetUserProgress(context.Context, *GetUserProgressRequest) (*GetUserProgressResponse, error)
	mustEmbedUnimplementedUserProgressServiceServer()
}

// UnimplementedUserProgressServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserProgressServiceServer struct {
}

func (UnimplementedUserProgressServiceServer) AddUserProgress(context.Context, *AddUserProgressRequest) (*AddUserProgressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserProgress not implemented")
}
func (UnimplementedUserProgressServiceServer) GetUserProgress(context.Context, *GetUserProgressRequest) (*GetUserProgressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProgress not implemented")
}
func (UnimplementedUserProgressServiceServer) mustEmbedUnimplementedUserProgressServiceServer() {}

// UnsafeUserProgressServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserProgressServiceServer will
// result in compilation errors.
type UnsafeUserProgressServiceServer interface {
	mustEmbedUnimplementedUserProgressServiceServer()
}

func RegisterUserProgressServiceServer(s grpc.ServiceRegistrar, srv UserProgressServiceServer) {
	s.RegisterService(&UserProgressService_ServiceDesc, srv)
}

func _UserProgressService_AddUserProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserProgressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProgressServiceServer).AddUserProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserProgressService_AddUserProgress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProgressServiceServer).AddUserProgress(ctx, req.(*AddUserProgressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProgressService_GetUserProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserProgressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProgressServiceServer).GetUserProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserProgressService_GetUserProgress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProgressServiceServer).GetUserProgress(ctx, req.(*GetUserProgressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserProgressService_ServiceDesc is the grpc.ServiceDesc for UserProgressService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserProgressService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "languagelearning.UserProgressService",
	HandlerType: (*UserProgressServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUserProgress",
			Handler:    _UserProgressService_AddUserProgress_Handler,
		},
		{
			MethodName: "GetUserProgress",
			Handler:    _UserProgressService_GetUserProgress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_progress.proto",
}