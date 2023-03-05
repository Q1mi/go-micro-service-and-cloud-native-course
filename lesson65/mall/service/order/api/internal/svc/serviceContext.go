package svc

import (
	"mall/service/order/api/internal/config"
	"mall/service/user/rpc/userclient" // RPC客户端代码

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRPC userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRPC: userclient.NewUser(zrpc.MustNewClient(c.UserRPC)),
	}
}
