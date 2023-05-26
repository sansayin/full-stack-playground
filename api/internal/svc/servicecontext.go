package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-rest/api/internal/config"
	"go-rest/rpc/user"
)

type ServiceContext struct {
	Config config.Config
	User   user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
