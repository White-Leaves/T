package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gomall/apps/order/cmd/rpc/internal/svc"
	"gomall/apps/order/cmd/rpc/pb"
	"gomall/pkg/snowflake"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {

	osn := snowflake.GenString()
	err := l.svcCtx.OrderModel.CreateOrder(l.ctx, osn, in.UserId, in.ProductId)
	if err != nil {
		logx.Errorf("OrderModel.CreateOrder osn:%v err:%v", osn, err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.CreateOrderResp{
		Sn: osn,
	}, nil

}
