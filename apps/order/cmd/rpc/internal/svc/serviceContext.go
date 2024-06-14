package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"gomall/apps/order/cmd/rpc/internal/config"
	"gomall/apps/order/model"
	"gomall/apps/product/cmd/rpc/product"
	"gomall/apps/user/cmd/rpc/user"
)

type ServiceContext struct {
	Config config.Config

	UserRpc    user.UserZrpcClient
	ProductRpc product.Product

	OrderModel     model.OrdersModel
	OrderitemModel model.OrderitemModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config: c,

		OrderModel:     model.NewOrdersModel(conn, c.Cache),
		OrderitemModel: model.NewOrderitemModel(conn, c.Cache),

		UserRpc:    user.NewUserZrpcClient(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc: product.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
