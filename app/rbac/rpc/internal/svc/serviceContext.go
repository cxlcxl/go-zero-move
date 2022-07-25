package svc

import (
	"business/app/rbac/rpc/internal/config"
	"business/app/rbac/rpc/model"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	RoleModel model.RolesModel
	UserModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	DBConn := sqlx.NewSqlConn(c.Database.Driver, getDsn(c))
	return &ServiceContext{
		Config:    c,
		RoleModel: model.NewRolesModel(DBConn, c.Cache),
		UserModel: model.NewUsersModel(DBConn, c.Cache),
	}
}

// mysql dsn 连接
func getDsn(c config.Config) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		c.Database.User,
		c.Database.Pass,
		c.Database.Host,
		c.Database.Port,
		c.Database.DbName,
		c.Database.Charset,
	)
}
