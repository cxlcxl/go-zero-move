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
	ApiModuleRefreshToken = "RefreshToken" // Token 刷新
	ApiModuleTargeting    = "Targeting"    // 定向列表
	ApiModulePosition     = "Position"     // 版位
)

var (
	SvcCtx *svc.ServiceContext
)
