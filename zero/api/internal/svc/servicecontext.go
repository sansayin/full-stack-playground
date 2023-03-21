package svc

import (
	"go-rest/zero/api/internal/config"
	user "go-rest/zero/rpc/user/users"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
  Users user.Users
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
    Users: user.NewUsers(zrpc.MustNewClient(c.Users)),
	}
}
