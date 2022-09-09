package svc

import (
	"business/app/rbac/rpc/internal/config"
	"business/app/rbac/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	RoleModel model.RolesModel
	UserModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	DBConn := sqlx.NewSqlConn(c.Database.Driver, c.Database.Dsn)
	return &ServiceContext{
		Config:    c,
		RoleModel: model.NewRolesModel(DBConn, c.Cache),
		UserModel: model.NewUsersModel(DBConn, c.Cache),
	}
}
