package svc

import (
	model2 "business/app/account/rpc/model"
	"business/app/marketing/rpc/internal/config"
	"business/common/utils"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	AccountModel model2.AccountsModel
	TokenModel   model2.TokensModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	DBConn := sqlx.NewSqlConn(c.Database.Driver, utils.GetDsn(c.Database.User, c.Database.Pass, c.Database.Host, c.Database.DbName, c.Database.Charset, c.Database.Port))
	return &ServiceContext{
		Config:       c,
		AccountModel: model2.NewAccountsModel(DBConn, c.Cache),
		TokenModel:   model2.NewTokensModel(DBConn, c.Cache),
	}
}
