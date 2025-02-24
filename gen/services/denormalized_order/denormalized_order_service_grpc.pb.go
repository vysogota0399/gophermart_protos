// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: denormalized_order_service.proto

package denormalized_order

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
	DenormalizedOrderService_OrderDetails_FullMethodName = "/DenormalizedOrderService/OrderDetails"
)

// DenormalizedOrderServiceClient is the client API for DenormalizedOrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DenormalizedOrderServiceClient interface {
	OrderDetails(ctx context.Context, in *OrderDetailsRequest, opts ...grpc.CallOption) (*DenormalizedOrder, error)
}

type denormalizedOrderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDenormalizedOrderServiceClient(cc grpc.ClientConnInterface) DenormalizedOrderServiceClient {
	return &denormalizedOrderServiceClient{cc}
}

func (c *denormalizedOrderServiceClient) OrderDetails(ctx context.Context, in *OrderDetailsRequest, opts ...grpc.CallOption) (*DenormalizedOrder, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DenormalizedOrder)
	err := c.cc.Invoke(ctx, DenormalizedOrderService_OrderDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DenormalizedOrderServiceServer is the server API for DenormalizedOrderService service.
// All implementations must embed UnimplementedDenormalizedOrderServiceServer
// for forward compatibility.
type DenormalizedOrderServiceServer interface {
	OrderDetails(context.Context, *OrderDetailsRequest) (*DenormalizedOrder, error)
	mustEmbedUnimplementedDenormalizedOrderServiceServer()
}

// UnimplementedDenormalizedOrderServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDenormalizedOrderServiceServer struct{}

func (UnimplementedDenormalizedOrderServiceServer) OrderDetails(context.Context, *OrderDetailsRequest) (*DenormalizedOrder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderDetails not implemented")
}
func (UnimplementedDenormalizedOrderServiceServer) mustEmbedUnimplementedDenormalizedOrderServiceServer() {
}
func (UnimplementedDenormalizedOrderServiceServer) testEmbeddedByValue() {}

// UnsafeDenormalizedOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DenormalizedOrderServiceServer will
// result in compilation errors.
type UnsafeDenormalizedOrderServiceServer interface {
	mustEmbedUnimplementedDenormalizedOrderServiceServer()
}

func RegisterDenormalizedOrderServiceServer(s grpc.ServiceRegistrar, srv DenormalizedOrderServiceServer) {
	// If the following call pancis, it indicates UnimplementedDenormalizedOrderServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DenormalizedOrderService_ServiceDesc, srv)
}

func _DenormalizedOrderService_OrderDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DenormalizedOrderServiceServer).OrderDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DenormalizedOrderService_OrderDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DenormalizedOrderServiceServer).OrderDetails(ctx, req.(*OrderDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DenormalizedOrderService_ServiceDesc is the grpc.ServiceDesc for DenormalizedOrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DenormalizedOrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DenormalizedOrderService",
	HandlerType: (*DenormalizedOrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OrderDetails",
			Handler:    _DenormalizedOrderService_OrderDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "denormalized_order_service.proto",
}
