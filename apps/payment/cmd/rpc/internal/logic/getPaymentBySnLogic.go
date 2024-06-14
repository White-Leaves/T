package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gomall/apps/payment/cmd/rpc/internal/svc"
	"gomall/apps/payment/cmd/rpc/pb"
	"gomall/apps/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPaymentBySnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPaymentBySnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPaymentBySnLogic {
	return &GetPaymentBySnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据sn查询流水记录
func (l *GetPaymentBySnLogic) GetPaymentBySn(in *pb.GetPaymentBySnReq) (*pb.GetPaymentBySnResp, error) {

	payment, err := l.svcCtx.PaymentModel.FindOneBySn(l.ctx, in.Sn)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(err, "GetPaymentBySn  FindOneBySn  db err:%v , in : %+v", err, in)
	}

	var resp pb.PaymentDetail
	if payment != nil {
		_ = copier.Copy(&resp, payment)

		resp.CreateTime = payment.CreateTime.Unix()
		resp.UpdateTime = payment.UpdateTime.Unix()
		resp.PayTime = payment.PayTime.Unix()
	}

	return &pb.GetPaymentBySnResp{
		PaymentDetail: &resp,
	}, nil
}
