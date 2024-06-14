package logic

import (
	"context"

	"gomall/apps/order/cmd/rpc/internal/svc"
	"gomall/apps/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrdersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrdersLogic {
	return &OrdersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrdersLogic) Orders(in *pb.OrdersReq) (*pb.OrdersResp, error) {

	return &pb.OrdersResp{}, nil
}
