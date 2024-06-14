package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gomall/apps/order/cmd/rpc/order"
	"gomall/apps/payment/cmd/api/internal/config"
	"gomall/apps/payment/cmd/rpc/payment"
	"gomall/apps/user/cmd/rpc/user"
)

type ServiceContext struct {
	Config config.Config

	PaymentRpc payment.Payment
	OrderRpc   order.Order
	UserRpc    user.UserZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		PaymentRpc: payment.NewPayment(zrpc.MustNewClient(c.PaymentRpcConf)),
		OrderRpc:   order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		UserRpc:    user.NewUserZrpcClient(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
