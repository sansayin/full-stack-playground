package svc

import (
	"go-rest/model"
	"go-rest/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
  ZeroUserModel model.ZeroUserModel  
}

func NewServiceContext(c config.Config) *ServiceContext {
  psqlConn:=sqlx.NewSqlConn("postgres", c.Psql.DataSource)
	return &ServiceContext{
		Config: c,
    ZeroUserModel:  model.NewZeroUserModel(psqlConn,c.CacheRedis),
	}
}
