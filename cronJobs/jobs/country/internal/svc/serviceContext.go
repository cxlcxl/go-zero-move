package svc

import (
	"business/common/utils"
	"business/cronJobs/jobs/country/config"
	model2 "business/cronJobs/jobs/country/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config             config.Config
	ReportCountryModel model2.ReportCountriesModel
	AdsLogModel        model2.AdsRequestLogsModel
	TokenModel         model2.TokensModel
	AccountModel       model2.AccountsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := utils.GetDsn(c.Database.User, c.Database.Pass, c.Database.Host, c.Database.DbName, c.Database.Charset, c.Database.Port)
	conn := sqlx.NewSqlConn(c.Database.Driver, dsn)
	return &ServiceContext{
		Config:             c,
		ReportCountryModel: model2.NewReportCountriesModel(conn),
		AdsLogModel:        model2.NewAdsRequestLogsModel(conn),
		TokenModel:         model2.NewTokensModel(conn),
		AccountModel:       model2.NewAccountsModel(conn),
	}
}
