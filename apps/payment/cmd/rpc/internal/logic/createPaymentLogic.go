package logic

import (
	"context"
	"github.com/pkg/errors"
	"gomall/apps/payment/model"
	"gomall/pkg/snowflake"

	"gomall/apps/payment/cmd/rpc/internal/svc"
	"gomall/apps/payment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePaymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePaymentLogic {
	return &CreatePaymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePaymentLogic) CreatePayment(in *pb.CreatePaymentReq) (*pb.CreatePaymentResp, error) {

	data := new(model.Payinfo)
	data.Sn = snowflake.GenString()
	data.Userid = in.Uid
	data.PayTotal = float64(in.PayTotal)
	data.OrderSn = in.OrderSn
	data.Platform = in.Platform
	data.PayStatus = 0
	_, err := l.svcCtx.PaymentModel.Insert(l.ctx, data)
	if err != nil {
		return nil, errors.Wrapf(err, "create new payment db insert fail,err:%v", err)

	}

	return &pb.CreatePaymentResp{
		Sn: data.Sn,
	}, nil
}
