// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: protos/learning/vocabulary.proto

package genprotos

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
	VocabularyService_AddVocabulary_FullMethodName    = "/languagelearning.VocabularyService/AddVocabulary"
	VocabularyService_GetVocabulary_FullMethodName    = "/languagelearning.VocabularyService/GetVocabulary"
	VocabularyService_UpdateVocabulary_FullMethodName = "/languagelearning.VocabularyService/UpdateVocabulary"
	VocabularyService_DeleteVocabulary_FullMethodName = "/languagelearning.VocabularyService/DeleteVocabulary"
)

// VocabularyServiceClient is the client API for VocabularyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Vocabulary xizmatini aniqlash
type VocabularyServiceClient interface {
	// Leksika yozuvini qo'shish uchun RPC
	AddVocabulary(ctx context.Context, in *AddVocabularyRequest, opts ...grpc.CallOption) (*AddVocabularyResponse, error)
	// Leksika yozuvini olish uchun RPC
	GetVocabulary(ctx context.Context, in *GetVocabularyRequest, opts ...grpc.CallOption) (*GetVocabularyResponse, error)
	// Leksika yozuvi yangilash uchun RPC
	UpdateVocabulary(ctx context.Context, in *UpdateVocabularyRequest, opts ...grpc.CallOption) (*UpdateVocabularyResponse, error)
	// Leksika yozuvi o'chirish uchun RPC
	DeleteVocabulary(ctx context.Context, in *DeleteVocabularyRequest, opts ...grpc.CallOption) (*DeleteVocabularyResponse, error)
}

type vocabularyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVocabularyServiceClient(cc grpc.ClientConnInterface) VocabularyServiceClient {
	return &vocabularyServiceClient{cc}
}

func (c *vocabularyServiceClient) AddVocabulary(ctx context.Context, in *AddVocabularyRequest, opts ...grpc.CallOption) (*AddVocabularyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddVocabularyResponse)
	err := c.cc.Invoke(ctx, VocabularyService_AddVocabulary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vocabularyServiceClient) GetVocabulary(ctx context.Context, in *GetVocabularyRequest, opts ...grpc.CallOption) (*GetVocabularyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVocabularyResponse)
	err := c.cc.Invoke(ctx, VocabularyService_GetVocabulary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vocabularyServiceClient) UpdateVocabulary(ctx context.Context, in *UpdateVocabularyRequest, opts ...grpc.CallOption) (*UpdateVocabularyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateVocabularyResponse)
	err := c.cc.Invoke(ctx, VocabularyService_UpdateVocabulary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vocabularyServiceClient) DeleteVocabulary(ctx context.Context, in *DeleteVocabularyRequest, opts ...grpc.CallOption) (*DeleteVocabularyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteVocabularyResponse)
	err := c.cc.Invoke(ctx, VocabularyService_DeleteVocabulary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VocabularyServiceServer is the server API for VocabularyService service.
// All implementations must embed UnimplementedVocabularyServiceServer
// for forward compatibility
//
// Vocabulary xizmatini aniqlash
type VocabularyServiceServer interface {
	// Leksika yozuvini qo'shish uchun RPC
	AddVocabulary(context.Context, *AddVocabularyRequest) (*AddVocabularyResponse, error)
	// Leksika yozuvini olish uchun RPC
	GetVocabulary(context.Context, *GetVocabularyRequest) (*GetVocabularyResponse, error)
	// Leksika yozuvi yangilash uchun RPC
	UpdateVocabulary(context.Context, *UpdateVocabularyRequest) (*UpdateVocabularyResponse, error)
	// Leksika yozuvi o'chirish uchun RPC
	DeleteVocabulary(context.Context, *DeleteVocabularyRequest) (*DeleteVocabularyResponse, error)
	mustEmbedUnimplementedVocabularyServiceServer()
}

// UnimplementedVocabularyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVocabularyServiceServer struct {
}

func (UnimplementedVocabularyServiceServer) AddVocabulary(context.Context, *AddVocabularyRequest) (*AddVocabularyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVocabulary not implemented")
}
func (UnimplementedVocabularyServiceServer) GetVocabulary(context.Context, *GetVocabularyRequest) (*GetVocabularyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVocabulary not implemented")
}
func (UnimplementedVocabularyServiceServer) UpdateVocabulary(context.Context, *UpdateVocabularyRequest) (*UpdateVocabularyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVocabulary not implemented")
}
func (UnimplementedVocabularyServiceServer) DeleteVocabulary(context.Context, *DeleteVocabularyRequest) (*DeleteVocabularyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVocabulary not implemented")
}
func (UnimplementedVocabularyServiceServer) mustEmbedUnimplementedVocabularyServiceServer() {}

// UnsafeVocabularyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VocabularyServiceServer will
// result in compilation errors.
type UnsafeVocabularyServiceServer interface {
	mustEmbedUnimplementedVocabularyServiceServer()
}

func RegisterVocabularyServiceServer(s grpc.ServiceRegistrar, srv VocabularyServiceServer) {
	s.RegisterService(&VocabularyService_ServiceDesc, srv)
}

func _VocabularyService_AddVocabulary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddVocabularyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VocabularyServiceServer).AddVocabulary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VocabularyService_AddVocabulary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VocabularyServiceServer).AddVocabulary(ctx, req.(*AddVocabularyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VocabularyService_GetVocabulary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVocabularyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VocabularyServiceServer).GetVocabulary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VocabularyService_GetVocabulary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VocabularyServiceServer).GetVocabulary(ctx, req.(*GetVocabularyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VocabularyService_UpdateVocabulary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVocabularyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VocabularyServiceServer).UpdateVocabulary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VocabularyService_UpdateVocabulary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VocabularyServiceServer).UpdateVocabulary(ctx, req.(*UpdateVocabularyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VocabularyService_DeleteVocabulary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVocabularyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VocabularyServiceServer).DeleteVocabulary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VocabularyService_DeleteVocabulary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VocabularyServiceServer).DeleteVocabulary(ctx, req.(*DeleteVocabularyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VocabularyService_ServiceDesc is the grpc.ServiceDesc for VocabularyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VocabularyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "languagelearning.VocabularyService",
	HandlerType: (*VocabularyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddVocabulary",
			Handler:    _VocabularyService_AddVocabulary_Handler,
		},
		{
			MethodName: "GetVocabulary",
			Handler:    _VocabularyService_GetVocabulary_Handler,
		},
		{
			MethodName: "UpdateVocabulary",
			Handler:    _VocabularyService_UpdateVocabulary_Handler,
		},
		{
			MethodName: "DeleteVocabulary",
			Handler:    _VocabularyService_DeleteVocabulary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/learning/vocabulary.proto",
}
