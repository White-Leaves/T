package logic

import (
	"context"
	"github.com/pkg/errors"
	"gomall/apps/payment/cmd/rpc/internal/svc"
	"gomall/apps/payment/cmd/rpc/pb"
	"gomall/apps/user/model"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTradeStateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTradeStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTradeStateLogic {
	return &UpdateTradeStateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新交易状态
func (l *UpdateTradeStateLogic) UpdateTradeState(in *pb.UpdateTradeStateReq) (*pb.UpdateTradeStateResp, error) {
	payment, err := l.svcCtx.PaymentModel.FindOneBySn(l.ctx, in.Sn)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(err, "Update Trade State FindBySn fail")
	}

	if payment == nil {
		return nil, errors.Wrapf(err, "payment record not exist")
	}

	payment.PayStatus = in.TradeState
	payment.PayTime = time.Now()

	err = l.svcCtx.PaymentModel.Update(l.ctx, payment)
	if err != nil {
		return nil, errors.Wrapf(err, "update Payment fail")
	}

	return &pb.UpdateTradeStateResp{}, nil
}
