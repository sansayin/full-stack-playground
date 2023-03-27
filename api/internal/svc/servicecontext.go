package svc

import (
	"go-rest/api/internal/config"
	"go-rest/rpc/user"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config 
  User user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
    User: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
