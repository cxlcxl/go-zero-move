package svc

import (
	"business/app/marketing/rpc/internal/config"
	"business/app/marketing/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	CampaignModel model.CampaignsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	DBConn := sqlx.NewSqlConn(c.Database.Driver, c.Database.Dsn)
	return &ServiceContext{
		Config:        c,
		CampaignModel: model.NewCampaignsModel(DBConn),
	}
}
