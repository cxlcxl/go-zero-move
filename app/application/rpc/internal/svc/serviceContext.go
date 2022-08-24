package svc

import (
	"business/app/application/rpc/internal/config"
	"business/app/application/rpc/model"
	"business/common/utils"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config   config.Config
	AppModel model.AppsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	DBConn := sqlx.NewSqlConn(c.Database.Driver, utils.GetDsn(c.Database.User, c.Database.Pass, c.Database.Host, c.Database.DbName, c.Database.Charset, c.Database.Port))
	return &ServiceContext{
		Config:   c,
		AppModel: model.NewAppsModel(DBConn, c.Cache),
	}
}
