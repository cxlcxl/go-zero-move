package svc

import (
	"business/app/marketing/rpc/internal/config"
	"business/app/marketing/rpc/model"
	"business/common/utils"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	AccountModel model.AccountsModel
	TokenModel   model.TokensModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	DBConn := sqlx.NewSqlConn(c.Database.Driver, utils.GetDsn(c.Database.User, c.Database.Pass, c.Database.Host, c.Database.DbName, c.Database.Charset, c.Database.Port))
	return &ServiceContext{
		Config:       c,
		AccountModel: model.NewAccountsModel(DBConn, c.Cache),
		TokenModel:   model.NewTokensModel(DBConn, c.Cache),
	}
}
