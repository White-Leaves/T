package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gomall/apps/user/cmd/api/internal/config"
	"gomall/apps/user/cmd/rpc/user"
)

type ServiceContext struct {
	Config config.Config

	UserRpc user.UserZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUserZrpcClient(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
