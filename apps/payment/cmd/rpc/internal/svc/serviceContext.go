package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gomall/apps/payment/cmd/rpc/internal/config"
	"gomall/apps/payment/model"
)

type ServiceContext struct {
	Config                  config.Config
	KqUpdatePayStatusClient *kq.Pusher
	PaymentModel            model.PayinfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:                  c,
		PaymentModel:            model.NewPayinfoModel(sqlConn, c.Cache),
		KqUpdatePayStatusClient: kq.NewPusher(c.KqPaymentUpdatePayStatusConf.Brokers, c.KqPaymentUpdatePayStatusConf.Topic),
	}
}
