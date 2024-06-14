package payment

import (
	"context"
	"github.com/pkg/errors"
	"gomall/apps/order/cmd/rpc/order"
	"gomall/apps/payment/cmd/api/internal/svc"
	"gomall/apps/payment/cmd/api/internal/types"
	"gomall/apps/payment/cmd/rpc/payment"
	"gomall/pkg/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentLogic {
	return &PaymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaymentLogic) Payment(req *types.PaymentReq) (*types.PaymentResp, error) {

	var totalPrice float64

	Price, err := l.getProductPrice(req.OrderSn)
	if err != nil {
		return nil, errors.Wrapf(err, "get product price err orderSn:%v", req.OrderSn)
	}

	totalPrice = Price

	userId := ctxdata.GetUidFromCtx(l.ctx)

	creatPaymentResp, err := l.svcCtx.PaymentRpc.CreatePayment(l.ctx, &payment.CreatePaymentReq{
		Uid:      userId,
		PayTotal: int64(totalPrice),
		OrderSn:  req.OrderSn,
		Platform: 1,
	})

	if err != nil || req.OrderSn == "" {
		return nil, errors.Wrapf(err, "create local payment fail")
	}

	_, err = l.svcCtx.PaymentRpc.UpdateTradeState(l.ctx, &payment.UpdateTradeStateReq{
		Sn:         creatPaymentResp.Sn,
		TradeState: 2,
	})

	return &types.PaymentResp{
		PaySn: creatPaymentResp.Sn,
	}, nil
}

func (l *PaymentLogic) getProductPrice(orderSn string) (float64, error) {

	resp, err := l.svcCtx.OrderRpc.GetOrderBySn(l.ctx, &order.GetOrderBySnReq{
		Sn: orderSn,
	})

	if err != nil {
		return 0, errors.Wrapf(err, "OrderRpc.HomestayOrderDetail err: %v, orderSn: %s", err, orderSn)

	}
	if resp.Order == nil || resp.Order.Id == 0 {
		return 0, errors.Wrapf(err, "Order not exist")
	}

	return resp.Order.Payment, nil
}
