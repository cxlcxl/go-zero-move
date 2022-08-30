package svc

import (
	"business/app/marketing/rpc/internal/config"
	"business/app/marketing/rpc/model"
	"business/common/utils"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	CampaignModel model.CampaignsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	DBConn := sqlx.NewSqlConn(c.Database.Driver, utils.GetDsn(c.Database.User, c.Database.Pass, c.Database.Host, c.Database.DbName, c.Database.Charset, c.Database.Port))
	return &ServiceContext{
		Config:        c,
		CampaignModel: model.NewCampaignsModel(DBConn),
	}
}
