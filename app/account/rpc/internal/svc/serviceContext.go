package svc

import (
	"business/app/account/rpc/internal/config"
	"business/app/account/rpc/model"
	"business/common/utils"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	ActModel       model.AccountsModel
	ActClientModel model.AccountClientIdsModel
	TokenModel     model.TokensModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	DBConn := sqlx.NewSqlConn(c.Database.Driver, utils.GetDsn(c.Database.User, c.Database.Pass, c.Database.Host, c.Database.DbName, c.Database.Charset, c.Database.Port))
	return &ServiceContext{
		Config:         c,
		ActModel:       model.NewAccountsModel(DBConn, c.Cache),
		ActClientModel: model.NewAccountClientIdsModel(DBConn, c.Cache),
		TokenModel:     model.NewTokensModel(DBConn, c.Cache),
	}
}
