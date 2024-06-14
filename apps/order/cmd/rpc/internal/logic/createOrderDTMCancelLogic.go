package logic

import (
	"context"
	"fmt"

	"gomall/apps/order/cmd/rpc/internal/svc"
	"gomall/apps/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderDTMCancelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderDTMCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderDTMCancelLogic {
	return &CreateOrderDTMCancelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderDTMCancelLogic) CreateOrderDTMCancel(in *pb.AddOrderReq) (*pb.AddOrderResp, error) {

	fmt.Printf("Create order cancel: %+v\n", in)
	return &pb.AddOrderResp{}, nil
}
