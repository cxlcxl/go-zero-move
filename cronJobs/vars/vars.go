package vars

import (
	"business/cronJobs/svc"
)

const (
	DateFormat = "2006-01-02"
	Env        = "dev"

	ApiModuleCountry      = "Country"
	ApiModuleCampaign     = "Campaign"
	ApiModuleDictionary   = "Dictionary"   // 定向字典数据
	ApiModuleRefreshToken = "RefreshToken" //
	ApiModuleTargeting    = "Targeting"    //
)

var (
	SvcCtx *svc.ServiceContext
)
