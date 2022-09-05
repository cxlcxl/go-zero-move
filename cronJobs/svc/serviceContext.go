package svc

import (
	"business/cronJobs/config"
	"business/cronJobs/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config             config.Config
	ReportCountryModel model.ReportCountriesModel
	JobModel           model.JobsModel
	TokenModel         model.TokensModel
	AccountModel       model.AccountsModel
	AppModel           model.AppsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn(c.Database.Driver, c.Database.Dsn)
	return &ServiceContext{
		Config:             c,
		ReportCountryModel: model.NewReportCountriesModel(conn),
		JobModel:           model.NewJobsModel(conn),
		TokenModel:         model.NewTokensModel(conn),
		AccountModel:       model.NewAccountsModel(conn),
		AppModel:           model.NewAppsModel(conn),
	}
}