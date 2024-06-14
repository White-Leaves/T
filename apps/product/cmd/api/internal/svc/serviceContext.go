package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"gomall/apps/product/cmd/api/internal/config"
	"gomall/apps/product/cmd/rpc/product"
	"gomall/apps/product/model"
)

type ServiceContext struct {
	Config config.Config

	ProductRpc product.Product

	ProductModel model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config:       c,
		ProductRpc:   product.NewProduct(zrpc.MustNewClient(c.ProductRpcConf)),
		ProductModel: model.NewProductModel(sqlConn, c.Cache),
	}
}
