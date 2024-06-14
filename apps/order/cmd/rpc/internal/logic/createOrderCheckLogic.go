package logic

import (
	"context"
	"fmt"

	"gomall/apps/order/cmd/rpc/internal/svc"
	"gomall/apps/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderCheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderCheckLogic {
	return &CreateOrderCheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderCheckLogic) CreateOrderCheck(in *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {
	fmt.Printf("Create Order check in:%v", in)

	return &pb.CreateOrderResp{}, nil
}
