// Code generated by goctl. DO NOT EDIT.
// Source: seckill.proto

package server

import (
	"context"

	"gomall/apps/seckill/cmd/rpc/internal/logic"
	"gomall/apps/seckill/cmd/rpc/internal/svc"
	"gomall/apps/seckill/cmd/rpc/pb"
)

type SeckillServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedSeckillServer
}

func NewSeckillServer(svcCtx *svc.ServiceContext) *SeckillServer {
	return &SeckillServer{
		svcCtx: svcCtx,
	}
}

func (s *SeckillServer) SeckillOrder(ctx context.Context, in *pb.SeckillOrderReq) (*pb.SeckillOrderResp, error) {
	l := logic.NewSeckillOrderLogic(ctx, s.svcCtx)
	return l.SeckillOrder(in)
}