package svc

import (
	"business/common/utils"
	"business/cronJobs/jobs/country/config"
	"business/cronJobs/jobs/country/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config             config.Config
	ReportCountryModel model.ReportCountriesModel
	AdsLogModel        model.AdsRequestLogsModel
	TokenModel         model.TokensModel
	AccountModel       model.AccountsModel
	AppModel           model.AppsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := utils.GetDsn(c.Database.User, c.Database.Pass, c.Database.Host, c.Database.DbName, c.Database.Charset, c.Database.Port)
	conn := sqlx.NewSqlConn(c.Database.Driver, dsn)
	return &ServiceContext{
		Config:             c,
		ReportCountryModel: model.NewReportCountriesModel(conn),
		AdsLogModel:        model.NewAdsRequestLogsModel(conn),
		TokenModel:         model.NewTokensModel(conn),
		AccountModel:       model.NewAccountsModel(conn),
		AppModel:           model.NewAppsModel(conn),
	}
}
