package svc

import (
	"go-rest/zero/rpc/user/internal/config"
	"go-rest/zero/rpc/user/model"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
  Model model.ZeroUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
  // conn:=sqlx.NewSqlConn("postgres", c.DataSource)
  // res, err := conn.Exec("select * from zero_user")
  // if(err!=nil){
  //   log.Println(err)
  // }
  // log.Println(res.RowsAffected())
	return &ServiceContext{
		Config: c,
   	Model: model.NewZeroUserModel(sqlx.NewSqlConn("postgres",c.DataSource), c.Cache),
	}
}
