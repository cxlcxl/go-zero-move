package svc

import (
	"business/app/application/rpc/internal/config"
	"business/app/application/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config   config.Config
	AppModel model.AppsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	DBConn := sqlx.NewSqlConn(c.Database.Driver, c.Database.Dsn)
	return &ServiceContext{
		Config:   c,
		AppModel: model.NewAppsModel(DBConn, c.Cache),
	}
}
