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
	DictModel          model.TargetingDictionariesModel
	TargetingModel     model.TargetingsModel
	PositionModel      model.PositionsModel
	PriceModel         model.PositionBasePricesModel
	ElementModel       model.PositionElementsModel
	CampaignModel      model.CampaignsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn(c.Database.Driver, c.Database.Dsn)
	return &ServiceContext{
		Config:             c,
		DictModel:          model.NewTargetingDictionariesModel(conn),
		ReportCountryModel: model.NewReportCountriesModel(conn),
		JobModel:           model.NewJobsModel(conn),
		TokenModel:         model.NewTokensModel(conn),
		AccountModel:       model.NewAccountsModel(conn),
		AppModel:           model.NewAppsModel(conn),
		TargetingModel:     model.NewTargetingsModel(conn),
		PositionModel:      model.NewPositionsModel(conn),
		PriceModel:         model.NewPositionBasePricesModel(conn),
		ElementModel:       model.NewPositionElementsModel(conn),
		CampaignModel:      model.NewCampaignsModel(conn),
	}
}
