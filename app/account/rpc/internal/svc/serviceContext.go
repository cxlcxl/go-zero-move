package svc

import (
	"business/app/account/rpc/internal/config"
	"business/app/account/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	AccountModel model.AccountsModel
	TokenModel   model.TokensModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	DBConn := sqlx.NewSqlConn(c.Database.Driver, c.Database.Dsn)
	return &ServiceContext{
		Config:       c,
		AccountModel: model.NewAccountsModel(DBConn, c.Cache),
		TokenModel:   model.NewTokensModel(DBConn),
	}
}
