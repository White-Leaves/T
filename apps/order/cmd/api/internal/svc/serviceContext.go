package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gomall/apps/order/cmd/api/internal/config"
	"gomall/apps/order/cmd/rpc/order"
	"gomall/apps/product/cmd/rpc/product"
)

type ServiceContext struct {
	Config config.Config

	OrderRpc order.Order

	ProductRpc product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		OrderRpc: order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),

		ProductRpc: product.NewProduct(zrpc.MustNewClient(c.ProductRpcConf)),
	}
}
