// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: delivery.proto

package models

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
	DeliveryService_UpdateDeliveryStatus_FullMethodName = "/models.DeliveryService/UpdateDeliveryStatus"
	DeliveryService_GetDelivery_FullMethodName          = "/models.DeliveryService/GetDelivery"
)

// DeliveryServiceClient is the client API for DeliveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeliveryServiceClient interface {
	UpdateDeliveryStatus(ctx context.Context, in *DeliveryStatus, opts ...grpc.CallOption) (*DeliveryResponse, error)
	GetDelivery(ctx context.Context, in *DeliveryID, opts ...grpc.CallOption) (*Delivery, error)
}

type deliveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeliveryServiceClient(cc grpc.ClientConnInterface) DeliveryServiceClient {
	return &deliveryServiceClient{cc}
}

func (c *deliveryServiceClient) UpdateDeliveryStatus(ctx context.Context, in *DeliveryStatus, opts ...grpc.CallOption) (*DeliveryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeliveryResponse)
	err := c.cc.Invoke(ctx, DeliveryService_UpdateDeliveryStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryServiceClient) GetDelivery(ctx context.Context, in *DeliveryID, opts ...grpc.CallOption) (*Delivery, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Delivery)
	err := c.cc.Invoke(ctx, DeliveryService_GetDelivery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeliveryServiceServer is the server API for DeliveryService service.
// All implementations must embed UnimplementedDeliveryServiceServer
// for forward compatibility.
type DeliveryServiceServer interface {
	UpdateDeliveryStatus(context.Context, *DeliveryStatus) (*DeliveryResponse, error)
	GetDelivery(context.Context, *DeliveryID) (*Delivery, error)
	mustEmbedUnimplementedDeliveryServiceServer()
}

// UnimplementedDeliveryServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDeliveryServiceServer struct{}

func (UnimplementedDeliveryServiceServer) UpdateDeliveryStatus(context.Context, *DeliveryStatus) (*DeliveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeliveryStatus not implemented")
}
func (UnimplementedDeliveryServiceServer) GetDelivery(context.Context, *DeliveryID) (*Delivery, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDelivery not implemented")
}
func (UnimplementedDeliveryServiceServer) mustEmbedUnimplementedDeliveryServiceServer() {}
func (UnimplementedDeliveryServiceServer) testEmbeddedByValue()                         {}

// UnsafeDeliveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeliveryServiceServer will
// result in compilation errors.
type UnsafeDeliveryServiceServer interface {
	mustEmbedUnimplementedDeliveryServiceServer()
}

func RegisterDeliveryServiceServer(s grpc.ServiceRegistrar, srv DeliveryServiceServer) {
	// If the following call pancis, it indicates UnimplementedDeliveryServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DeliveryService_ServiceDesc, srv)
}

func _DeliveryService_UpdateDeliveryStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliveryStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryServiceServer).UpdateDeliveryStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeliveryService_UpdateDeliveryStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryServiceServer).UpdateDeliveryStatus(ctx, req.(*DeliveryStatus))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveryService_GetDelivery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliveryID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryServiceServer).GetDelivery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeliveryService_GetDelivery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryServiceServer).GetDelivery(ctx, req.(*DeliveryID))
	}
	return interceptor(ctx, in, info, handler)
}

// DeliveryService_ServiceDesc is the grpc.ServiceDesc for DeliveryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeliveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "models.DeliveryService",
	HandlerType: (*DeliveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateDeliveryStatus",
			Handler:    _DeliveryService_UpdateDeliveryStatus_Handler,
		},
		{
			MethodName: "GetDelivery",
			Handler:    _DeliveryService_GetDelivery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "delivery.proto",
}
