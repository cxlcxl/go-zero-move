package vars

import (
	"business/cronJobs/svc"
)

const (
	DateFormat = "2006-01-02"
	Env        = "dev"

	ApiModuleCountry  = "Country"
	ApiModuleCampaign = "Campaign"
)

var (
	SvcCtx *svc.ServiceContext
)
