package svc

import (
	"business/app/rbac/rpc/internal/config"
	"business/app/rbac/rpc/model"
	"business/common/utils"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	RoleModel model.RolesModel
	UserModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	DBConn := sqlx.NewSqlConn(c.Database.Driver, utils.GetDsn(c.Database.User, c.Database.Pass, c.Database.Host, c.Database.DbName, c.Database.Charset, c.Database.Port))
	return &ServiceContext{
		Config:    c,
		RoleModel: model.NewRolesModel(DBConn, c.Cache),
		UserModel: model.NewUsersModel(DBConn, c.Cache),
	}
}
