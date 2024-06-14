package logic

import (
	"context"

	"gomall/apps/payment/cmd/rpc/internal/svc"
	"gomall/apps/payment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPaymentSuccessRefundByOrderSnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPaymentSuccessRefundByOrderSnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPaymentSuccessRefundByOrderSnLogic {
	return &GetPaymentSuccessRefundByOrderSnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据订单sn查询流水记录
func (l *GetPaymentSuccessRefundByOrderSnLogic) GetPaymentSuccessRefundByOrderSn(in *pb.GetPaymentSuccessRefundByOrderSnReq) (*pb.GetPaymentSuccessRefundByOrderSnResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetPaymentSuccessRefundByOrderSnResp{}, nil
}
