// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: proto/reserve/reserve.proto

package yuemnoi_reserve

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
	ItemService_CreateItem_FullMethodName = "/ItemService/CreateItem"
)

// ItemServiceClient is the client API for ItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The item service definition.
type ItemServiceClient interface {
	CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error)
}

type itemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewItemServiceClient(cc grpc.ClientConnInterface) ItemServiceClient {
	return &itemServiceClient{cc}
}

func (c *itemServiceClient) CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateItemResponse)
	err := c.cc.Invoke(ctx, ItemService_CreateItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItemServiceServer is the server API for ItemService service.
// All implementations must embed UnimplementedItemServiceServer
// for forward compatibility.
//
// The item service definition.
type ItemServiceServer interface {
	CreateItem(context.Context, *CreateItemRequest) (*CreateItemResponse, error)
	mustEmbedUnimplementedItemServiceServer()
}

// UnimplementedItemServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedItemServiceServer struct{}

func (UnimplementedItemServiceServer) CreateItem(context.Context, *CreateItemRequest) (*CreateItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateItem not implemented")
}
func (UnimplementedItemServiceServer) mustEmbedUnimplementedItemServiceServer() {}
func (UnimplementedItemServiceServer) testEmbeddedByValue()                     {}

// UnsafeItemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ItemServiceServer will
// result in compilation errors.
type UnsafeItemServiceServer interface {
	mustEmbedUnimplementedItemServiceServer()
}

func RegisterItemServiceServer(s grpc.ServiceRegistrar, srv ItemServiceServer) {
	// If the following call pancis, it indicates UnimplementedItemServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ItemService_ServiceDesc, srv)
}

func _ItemService_CreateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).CreateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItemService_CreateItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).CreateItem(ctx, req.(*CreateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ItemService_ServiceDesc is the grpc.ServiceDesc for ItemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ItemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ItemService",
	HandlerType: (*ItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateItem",
			Handler:    _ItemService_CreateItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/reserve/reserve.proto",
}

const (
	ReserveService_CreateLendingRequest_FullMethodName    = "/ReserveService/CreateLendingRequest"
	ReserveService_GetLendingRequestDetail_FullMethodName = "/ReserveService/GetLendingRequestDetail"
	ReserveService_RejectLendingRequest_FullMethodName    = "/ReserveService/RejectLendingRequest"
	ReserveService_AcceptLendingRequest_FullMethodName    = "/ReserveService/AcceptLendingRequest"
)

// ReserveServiceClient is the client API for ReserveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReserveServiceClient interface {
	CreateLendingRequest(ctx context.Context, in *CreateLendingRequestRequest, opts ...grpc.CallOption) (*CreateLendingRequestResponse, error)
	GetLendingRequestDetail(ctx context.Context, in *GetLendingRequestDetailRequest, opts ...grpc.CallOption) (*Request, error)
	RejectLendingRequest(ctx context.Context, in *RejectLendingRequestRequest, opts ...grpc.CallOption) (*Request, error)
	AcceptLendingRequest(ctx context.Context, in *AcceptLendingRequestRequest, opts ...grpc.CallOption) (*Request, error)
}

type reserveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReserveServiceClient(cc grpc.ClientConnInterface) ReserveServiceClient {
	return &reserveServiceClient{cc}
}

func (c *reserveServiceClient) CreateLendingRequest(ctx context.Context, in *CreateLendingRequestRequest, opts ...grpc.CallOption) (*CreateLendingRequestResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateLendingRequestResponse)
	err := c.cc.Invoke(ctx, ReserveService_CreateLendingRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reserveServiceClient) GetLendingRequestDetail(ctx context.Context, in *GetLendingRequestDetailRequest, opts ...grpc.CallOption) (*Request, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Request)
	err := c.cc.Invoke(ctx, ReserveService_GetLendingRequestDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reserveServiceClient) RejectLendingRequest(ctx context.Context, in *RejectLendingRequestRequest, opts ...grpc.CallOption) (*Request, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Request)
	err := c.cc.Invoke(ctx, ReserveService_RejectLendingRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reserveServiceClient) AcceptLendingRequest(ctx context.Context, in *AcceptLendingRequestRequest, opts ...grpc.CallOption) (*Request, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Request)
	err := c.cc.Invoke(ctx, ReserveService_AcceptLendingRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReserveServiceServer is the server API for ReserveService service.
// All implementations must embed UnimplementedReserveServiceServer
// for forward compatibility.
type ReserveServiceServer interface {
	CreateLendingRequest(context.Context, *CreateLendingRequestRequest) (*CreateLendingRequestResponse, error)
	GetLendingRequestDetail(context.Context, *GetLendingRequestDetailRequest) (*Request, error)
	RejectLendingRequest(context.Context, *RejectLendingRequestRequest) (*Request, error)
	AcceptLendingRequest(context.Context, *AcceptLendingRequestRequest) (*Request, error)
	mustEmbedUnimplementedReserveServiceServer()
}

// UnimplementedReserveServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedReserveServiceServer struct{}

func (UnimplementedReserveServiceServer) CreateLendingRequest(context.Context, *CreateLendingRequestRequest) (*CreateLendingRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLendingRequest not implemented")
}
func (UnimplementedReserveServiceServer) GetLendingRequestDetail(context.Context, *GetLendingRequestDetailRequest) (*Request, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLendingRequestDetail not implemented")
}
func (UnimplementedReserveServiceServer) RejectLendingRequest(context.Context, *RejectLendingRequestRequest) (*Request, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectLendingRequest not implemented")
}
func (UnimplementedReserveServiceServer) AcceptLendingRequest(context.Context, *AcceptLendingRequestRequest) (*Request, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptLendingRequest not implemented")
}
func (UnimplementedReserveServiceServer) mustEmbedUnimplementedReserveServiceServer() {}
func (UnimplementedReserveServiceServer) testEmbeddedByValue()                        {}

// UnsafeReserveServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReserveServiceServer will
// result in compilation errors.
type UnsafeReserveServiceServer interface {
	mustEmbedUnimplementedReserveServiceServer()
}

func RegisterReserveServiceServer(s grpc.ServiceRegistrar, srv ReserveServiceServer) {
	// If the following call pancis, it indicates UnimplementedReserveServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ReserveService_ServiceDesc, srv)
}

func _ReserveService_CreateLendingRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLendingRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReserveServiceServer).CreateLendingRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReserveService_CreateLendingRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReserveServiceServer).CreateLendingRequest(ctx, req.(*CreateLendingRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReserveService_GetLendingRequestDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLendingRequestDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReserveServiceServer).GetLendingRequestDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReserveService_GetLendingRequestDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReserveServiceServer).GetLendingRequestDetail(ctx, req.(*GetLendingRequestDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReserveService_RejectLendingRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectLendingRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReserveServiceServer).RejectLendingRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReserveService_RejectLendingRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReserveServiceServer).RejectLendingRequest(ctx, req.(*RejectLendingRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReserveService_AcceptLendingRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptLendingRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReserveServiceServer).AcceptLendingRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReserveService_AcceptLendingRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReserveServiceServer).AcceptLendingRequest(ctx, req.(*AcceptLendingRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReserveService_ServiceDesc is the grpc.ServiceDesc for ReserveService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReserveService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ReserveService",
	HandlerType: (*ReserveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLendingRequest",
			Handler:    _ReserveService_CreateLendingRequest_Handler,
		},
		{
			MethodName: "GetLendingRequestDetail",
			Handler:    _ReserveService_GetLendingRequestDetail_Handler,
		},
		{
			MethodName: "RejectLendingRequest",
			Handler:    _ReserveService_RejectLendingRequest_Handler,
		},
		{
			MethodName: "AcceptLendingRequest",
			Handler:    _ReserveService_AcceptLendingRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/reserve/reserve.proto",
}
