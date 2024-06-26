// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: seckill.proto

package pb

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
	Seckill_SeckillOrder_FullMethodName = "/seckill.Seckill/SeckillOrder"
)

// SeckillClient is the client API for Seckill service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SeckillClient interface {
	SeckillOrder(ctx context.Context, in *SeckillOrderReq, opts ...grpc.CallOption) (*SeckillOrderResp, error)
}

type seckillClient struct {
	cc grpc.ClientConnInterface
}

func NewSeckillClient(cc grpc.ClientConnInterface) SeckillClient {
	return &seckillClient{cc}
}

func (c *seckillClient) SeckillOrder(ctx context.Context, in *SeckillOrderReq, opts ...grpc.CallOption) (*SeckillOrderResp, error) {
	out := new(SeckillOrderResp)
	err := c.cc.Invoke(ctx, Seckill_SeckillOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SeckillServer is the server API for Seckill service.
// All implementations must embed UnimplementedSeckillServer
// for forward compatibility
type SeckillServer interface {
	SeckillOrder(context.Context, *SeckillOrderReq) (*SeckillOrderResp, error)
	mustEmbedUnimplementedSeckillServer()
}

// UnimplementedSeckillServer must be embedded to have forward compatible implementations.
type UnimplementedSeckillServer struct {
}

func (UnimplementedSeckillServer) SeckillOrder(context.Context, *SeckillOrderReq) (*SeckillOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SeckillOrder not implemented")
}
func (UnimplementedSeckillServer) mustEmbedUnimplementedSeckillServer() {}

// UnsafeSeckillServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SeckillServer will
// result in compilation errors.
type UnsafeSeckillServer interface {
	mustEmbedUnimplementedSeckillServer()
}

func RegisterSeckillServer(s grpc.ServiceRegistrar, srv SeckillServer) {
	s.RegisterService(&Seckill_ServiceDesc, srv)
}

func _Seckill_SeckillOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeckillOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeckillServer).SeckillOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Seckill_SeckillOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeckillServer).SeckillOrder(ctx, req.(*SeckillOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Seckill_ServiceDesc is the grpc.ServiceDesc for Seckill service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Seckill_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "seckill.Seckill",
	HandlerType: (*SeckillServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SeckillOrder",
			Handler:    _Seckill_SeckillOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "seckill.proto",
}
