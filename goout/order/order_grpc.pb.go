// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: order/order.proto

package order

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
	OrderSrv_SubmitOrder_FullMethodName       = "/order.OrderSrv/SubmitOrder"
	OrderSrv_UpdateOrderStatus_FullMethodName = "/order.OrderSrv/UpdateOrderStatus"
	OrderSrv_GetOrderDetail_FullMethodName    = "/order.OrderSrv/GetOrderDetail"
	OrderSrv_GetOrderList_FullMethodName      = "/order.OrderSrv/GetOrderList"
)

// OrderSrvClient is the client API for OrderSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderSrvClient interface {
	// 提交订单
	SubmitOrder(ctx context.Context, in *SubmitOrderRequest, opts ...grpc.CallOption) (*SubmitOrderResponse, error)
	// 更新订单状态
	UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusRequest, opts ...grpc.CallOption) (*UpdateOrderStatusResponse, error)
	// 查询订单详情
	GetOrderDetail(ctx context.Context, in *GetOrderDetailRequest, opts ...grpc.CallOption) (*GetOrderDetailResponse, error)
	// 查询订单列表
	GetOrderList(ctx context.Context, in *GetOrderListRequest, opts ...grpc.CallOption) (*GetOrderListResponse, error)
}

type orderSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderSrvClient(cc grpc.ClientConnInterface) OrderSrvClient {
	return &orderSrvClient{cc}
}

func (c *orderSrvClient) SubmitOrder(ctx context.Context, in *SubmitOrderRequest, opts ...grpc.CallOption) (*SubmitOrderResponse, error) {
	out := new(SubmitOrderResponse)
	err := c.cc.Invoke(ctx, OrderSrv_SubmitOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderSrvClient) UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusRequest, opts ...grpc.CallOption) (*UpdateOrderStatusResponse, error) {
	out := new(UpdateOrderStatusResponse)
	err := c.cc.Invoke(ctx, OrderSrv_UpdateOrderStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderSrvClient) GetOrderDetail(ctx context.Context, in *GetOrderDetailRequest, opts ...grpc.CallOption) (*GetOrderDetailResponse, error) {
	out := new(GetOrderDetailResponse)
	err := c.cc.Invoke(ctx, OrderSrv_GetOrderDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderSrvClient) GetOrderList(ctx context.Context, in *GetOrderListRequest, opts ...grpc.CallOption) (*GetOrderListResponse, error) {
	out := new(GetOrderListResponse)
	err := c.cc.Invoke(ctx, OrderSrv_GetOrderList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderSrvServer is the server API for OrderSrv service.
// All implementations must embed UnimplementedOrderSrvServer
// for forward compatibility
type OrderSrvServer interface {
	// 提交订单
	SubmitOrder(context.Context, *SubmitOrderRequest) (*SubmitOrderResponse, error)
	// 更新订单状态
	UpdateOrderStatus(context.Context, *UpdateOrderStatusRequest) (*UpdateOrderStatusResponse, error)
	// 查询订单详情
	GetOrderDetail(context.Context, *GetOrderDetailRequest) (*GetOrderDetailResponse, error)
	// 查询订单列表
	GetOrderList(context.Context, *GetOrderListRequest) (*GetOrderListResponse, error)
	mustEmbedUnimplementedOrderSrvServer()
}

// UnimplementedOrderSrvServer must be embedded to have forward compatible implementations.
type UnimplementedOrderSrvServer struct {
}

func (UnimplementedOrderSrvServer) SubmitOrder(context.Context, *SubmitOrderRequest) (*SubmitOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitOrder not implemented")
}
func (UnimplementedOrderSrvServer) UpdateOrderStatus(context.Context, *UpdateOrderStatusRequest) (*UpdateOrderStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderStatus not implemented")
}
func (UnimplementedOrderSrvServer) GetOrderDetail(context.Context, *GetOrderDetailRequest) (*GetOrderDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderDetail not implemented")
}
func (UnimplementedOrderSrvServer) GetOrderList(context.Context, *GetOrderListRequest) (*GetOrderListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderList not implemented")
}
func (UnimplementedOrderSrvServer) mustEmbedUnimplementedOrderSrvServer() {}

// UnsafeOrderSrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderSrvServer will
// result in compilation errors.
type UnsafeOrderSrvServer interface {
	mustEmbedUnimplementedOrderSrvServer()
}

func RegisterOrderSrvServer(s grpc.ServiceRegistrar, srv OrderSrvServer) {
	s.RegisterService(&OrderSrv_ServiceDesc, srv)
}

func _OrderSrv_SubmitOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderSrvServer).SubmitOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderSrv_SubmitOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderSrvServer).SubmitOrder(ctx, req.(*SubmitOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderSrv_UpdateOrderStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderSrvServer).UpdateOrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderSrv_UpdateOrderStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderSrvServer).UpdateOrderStatus(ctx, req.(*UpdateOrderStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderSrv_GetOrderDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderSrvServer).GetOrderDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderSrv_GetOrderDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderSrvServer).GetOrderDetail(ctx, req.(*GetOrderDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderSrv_GetOrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderSrvServer).GetOrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderSrv_GetOrderList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderSrvServer).GetOrderList(ctx, req.(*GetOrderListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderSrv_ServiceDesc is the grpc.ServiceDesc for OrderSrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderSrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.OrderSrv",
	HandlerType: (*OrderSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitOrder",
			Handler:    _OrderSrv_SubmitOrder_Handler,
		},
		{
			MethodName: "UpdateOrderStatus",
			Handler:    _OrderSrv_UpdateOrderStatus_Handler,
		},
		{
			MethodName: "GetOrderDetail",
			Handler:    _OrderSrv_GetOrderDetail_Handler,
		},
		{
			MethodName: "GetOrderList",
			Handler:    _OrderSrv_GetOrderList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order/order.proto",
}
