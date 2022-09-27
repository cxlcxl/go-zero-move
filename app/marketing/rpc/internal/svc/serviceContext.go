package svc

import (
	"business/app/marketing/rpc/internal/config"
	"business/app/marketing/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	CampaignModel  model.CampaignsModel
	DictModel      model.TargetingDictionariesModel
	RegionModel    model.OverseasRegionsModel
	ContinentModel model.ContinentsModel
	TargetingModel model.TargetingsModel
	PositionModel  model.PositionsModel
	SampleModel    model.PositionSamplesModel
	PlacementModel model.PositionPlacementsModel
	AssetModel     model.AssetsModel
	TrackModel     model.TrackingsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	DBConn := sqlx.NewSqlConn(c.Database.Driver, c.Database.Dsn)
	return &ServiceContext{
		Config:         c,
		CampaignModel:  model.NewCampaignsModel(DBConn),
		DictModel:      model.NewTargetingDictionariesModel(DBConn),
		RegionModel:    model.NewOverseasRegionsModel(DBConn),
		ContinentModel: model.NewContinentsModel(DBConn),
		TargetingModel: model.NewTargetingsModel(DBConn),
		PositionModel:  model.NewPositionsModel(DBConn),
		SampleModel:    model.NewPositionSamplesModel(DBConn),
		PlacementModel: model.NewPositionPlacementsModel(DBConn),
		AssetModel:     model.NewAssetsModel(DBConn),
		TrackModel:     model.NewTrackingsModel(DBConn),
	}
}
