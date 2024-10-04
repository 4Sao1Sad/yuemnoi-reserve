// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
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