package svc

import (
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-rest/model"
	"go-rest/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Model  model.ZeroUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewZeroUserModel(sqlx.NewSqlConn("postgres", c.Psql.DataSource), c.CacheRedis),
	}
}
