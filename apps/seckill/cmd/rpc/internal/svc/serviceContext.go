package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gomall/apps/product/cmd/rpc/product"
	"gomall/apps/seckill/cmd/rpc/internal/config"

	"github.com/zeromicro/go-queue/kq"
)

type ServiceContext struct {
	Config      config.Config
	BizRedis    *redis.Redis
	ProductRpc  product.Product
	KafkaPusher *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		BizRedis:    redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
		ProductRpc:  product.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
		KafkaPusher: kq.NewPusher(c.Kafka.Addrs, c.Kafka.Topic),
	}
}
