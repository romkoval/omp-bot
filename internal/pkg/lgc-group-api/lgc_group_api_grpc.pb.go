// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package lgc_group_api

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

// OmpGroupApiServiceClient is the client API for OmpGroupApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OmpGroupApiServiceClient interface {
	DescribeGroupV1(ctx context.Context, in *DescribeGroupV1Request, opts ...grpc.CallOption) (*DescribeGroupV1Response, error)
	CreateGroupV1(ctx context.Context, in *CreateGroupV1Request, opts ...grpc.CallOption) (*CreateGroupV1Response, error)
	UpdateGroupV1(ctx context.Context, in *UpdateGroupV1Request, opts ...grpc.CallOption) (*UpdateGroupV1Response, error)
	ListGroupV1(ctx context.Context, in *ListGroupV1Request, opts ...grpc.CallOption) (*ListGroupV1Response, error)
	RemoveGroupV1(ctx context.Context, in *RemoveGroupV1Request, opts ...grpc.CallOption) (*RemoveGroupV1Response, error)
}

type ompGroupApiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOmpGroupApiServiceClient(cc grpc.ClientConnInterface) OmpGroupApiServiceClient {
	return &ompGroupApiServiceClient{cc}
}

func (c *ompGroupApiServiceClient) DescribeGroupV1(ctx context.Context, in *DescribeGroupV1Request, opts ...grpc.CallOption) (*DescribeGroupV1Response, error) {
	out := new(DescribeGroupV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.lgc_group_api.v1.OmpGroupApiService/DescribeGroupV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ompGroupApiServiceClient) CreateGroupV1(ctx context.Context, in *CreateGroupV1Request, opts ...grpc.CallOption) (*CreateGroupV1Response, error) {
	out := new(CreateGroupV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.lgc_group_api.v1.OmpGroupApiService/CreateGroupV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ompGroupApiServiceClient) UpdateGroupV1(ctx context.Context, in *UpdateGroupV1Request, opts ...grpc.CallOption) (*UpdateGroupV1Response, error) {
	out := new(UpdateGroupV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.lgc_group_api.v1.OmpGroupApiService/UpdateGroupV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ompGroupApiServiceClient) ListGroupV1(ctx context.Context, in *ListGroupV1Request, opts ...grpc.CallOption) (*ListGroupV1Response, error) {
	out := new(ListGroupV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.lgc_group_api.v1.OmpGroupApiService/ListGroupV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ompGroupApiServiceClient) RemoveGroupV1(ctx context.Context, in *RemoveGroupV1Request, opts ...grpc.CallOption) (*RemoveGroupV1Response, error) {
	out := new(RemoveGroupV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.lgc_group_api.v1.OmpGroupApiService/RemoveGroupV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OmpGroupApiServiceServer is the server API for OmpGroupApiService service.
// All implementations must embed UnimplementedOmpGroupApiServiceServer
// for forward compatibility
type OmpGroupApiServiceServer interface {
	DescribeGroupV1(context.Context, *DescribeGroupV1Request) (*DescribeGroupV1Response, error)
	CreateGroupV1(context.Context, *CreateGroupV1Request) (*CreateGroupV1Response, error)
	UpdateGroupV1(context.Context, *UpdateGroupV1Request) (*UpdateGroupV1Response, error)
	ListGroupV1(context.Context, *ListGroupV1Request) (*ListGroupV1Response, error)
	RemoveGroupV1(context.Context, *RemoveGroupV1Request) (*RemoveGroupV1Response, error)
	mustEmbedUnimplementedOmpGroupApiServiceServer()
}

// UnimplementedOmpGroupApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOmpGroupApiServiceServer struct {
}

func (UnimplementedOmpGroupApiServiceServer) DescribeGroupV1(context.Context, *DescribeGroupV1Request) (*DescribeGroupV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeGroupV1 not implemented")
}
func (UnimplementedOmpGroupApiServiceServer) CreateGroupV1(context.Context, *CreateGroupV1Request) (*CreateGroupV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroupV1 not implemented")
}
func (UnimplementedOmpGroupApiServiceServer) UpdateGroupV1(context.Context, *UpdateGroupV1Request) (*UpdateGroupV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupV1 not implemented")
}
func (UnimplementedOmpGroupApiServiceServer) ListGroupV1(context.Context, *ListGroupV1Request) (*ListGroupV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroupV1 not implemented")
}
func (UnimplementedOmpGroupApiServiceServer) RemoveGroupV1(context.Context, *RemoveGroupV1Request) (*RemoveGroupV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveGroupV1 not implemented")
}
func (UnimplementedOmpGroupApiServiceServer) mustEmbedUnimplementedOmpGroupApiServiceServer() {}

// UnsafeOmpGroupApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OmpGroupApiServiceServer will
// result in compilation errors.
type UnsafeOmpGroupApiServiceServer interface {
	mustEmbedUnimplementedOmpGroupApiServiceServer()
}

func RegisterOmpGroupApiServiceServer(s grpc.ServiceRegistrar, srv OmpGroupApiServiceServer) {
	s.RegisterService(&OmpGroupApiService_ServiceDesc, srv)
}

func _OmpGroupApiService_DescribeGroupV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeGroupV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmpGroupApiServiceServer).DescribeGroupV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.lgc_group_api.v1.OmpGroupApiService/DescribeGroupV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmpGroupApiServiceServer).DescribeGroupV1(ctx, req.(*DescribeGroupV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OmpGroupApiService_CreateGroupV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmpGroupApiServiceServer).CreateGroupV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.lgc_group_api.v1.OmpGroupApiService/CreateGroupV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmpGroupApiServiceServer).CreateGroupV1(ctx, req.(*CreateGroupV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OmpGroupApiService_UpdateGroupV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmpGroupApiServiceServer).UpdateGroupV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.lgc_group_api.v1.OmpGroupApiService/UpdateGroupV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmpGroupApiServiceServer).UpdateGroupV1(ctx, req.(*UpdateGroupV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OmpGroupApiService_ListGroupV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmpGroupApiServiceServer).ListGroupV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.lgc_group_api.v1.OmpGroupApiService/ListGroupV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmpGroupApiServiceServer).ListGroupV1(ctx, req.(*ListGroupV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OmpGroupApiService_RemoveGroupV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveGroupV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmpGroupApiServiceServer).RemoveGroupV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.lgc_group_api.v1.OmpGroupApiService/RemoveGroupV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmpGroupApiServiceServer).RemoveGroupV1(ctx, req.(*RemoveGroupV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// OmpGroupApiService_ServiceDesc is the grpc.ServiceDesc for OmpGroupApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OmpGroupApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ozonmp.lgc_group_api.v1.OmpGroupApiService",
	HandlerType: (*OmpGroupApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeGroupV1",
			Handler:    _OmpGroupApiService_DescribeGroupV1_Handler,
		},
		{
			MethodName: "CreateGroupV1",
			Handler:    _OmpGroupApiService_CreateGroupV1_Handler,
		},
		{
			MethodName: "UpdateGroupV1",
			Handler:    _OmpGroupApiService_UpdateGroupV1_Handler,
		},
		{
			MethodName: "ListGroupV1",
			Handler:    _OmpGroupApiService_ListGroupV1_Handler,
		},
		{
			MethodName: "RemoveGroupV1",
			Handler:    _OmpGroupApiService_RemoveGroupV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ozonmp/lgc_group_api/v1/lgc_group_api.proto",
}
