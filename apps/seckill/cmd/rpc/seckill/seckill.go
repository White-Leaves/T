// Code generated by goctl. DO NOT EDIT.
// Source: seckill.proto

package seckill

import (
	"context"

	"gomall/apps/seckill/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	SeckillOrderReq  = pb.SeckillOrderReq
	SeckillOrderResp = pb.SeckillOrderResp

	Seckill interface {
		SeckillOrder(ctx context.Context, in *SeckillOrderReq, opts ...grpc.CallOption) (*SeckillOrderResp, error)
	}

	defaultSeckill struct {
		cli zrpc.Client
	}
)

func NewSeckill(cli zrpc.Client) Seckill {
	return &defaultSeckill{
		cli: cli,
	}
}

func (m *defaultSeckill) SeckillOrder(ctx context.Context, in *SeckillOrderReq, opts ...grpc.CallOption) (*SeckillOrderResp, error) {
	client := pb.NewSeckillClient(m.cli.Conn())
	return client.SeckillOrder(ctx, in, opts...)
}
