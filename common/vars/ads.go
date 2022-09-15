package vars

type AdsAccessTokenResponse struct {
	Error            int64  `json:"error"`
	ErrorDescription string `json:"error_description"`
	AccessToken      string `json:"access_token"`
	ExpiresIn        int64  `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
}

const (
	OptStatusEnable              = "OPERATION_STATUS_ENABLE"
	OptStatusDisable             = "OPERATION_STATUS_DISABLE"
	OptStatusDelete              = "OPERATION_STATUS_DELETE"
	ProductTypeWEB               = "WEB"
	ProductTypeAndroidApp        = "ANDROID_APP"
	CampaignTypeDisplay          = "CAMPAIGN_TYPE_DISPLAY"
	CampaignTypeShopping         = "CAMPAIGN_TYPE_SHOPPING"
	SyncFlowYes                  = "YES"
	SyncFlowNo                   = "NO"
	DailyBudgetUpdateToday       = "UPDATE_TODAY_DAILY_BUDGET"
	DailyBudgetUpdateTomorrow    = "UPDATE_TOMORROW_DAILY_BUDGET"
	DailyBudgetDeleteTomorrow    = "DELETE_TOMORROW_DAILY_BUDGET"
	CampaignStatusDelete         = "CAMPAIGN_STATUS_DELETE"
	CampaignStatusDisable        = "CAMPAIGN_STATUS_DISABLE"
	CampaignStatusFrozen         = "CAMPAIGN_STATUS_ADVERTISER_FROZEN"
	CampaignStatusBalance        = "CAMPAIGN_STATUS_ADVERTISER_BALANCE_EXCEED"
	CampaignStatusDailyBudget    = "CAMPAIGN_STATUS_ADVERTISER_DAILY_BUDGET_EXCEED"
	CampaignStatusDaily          = "CAMPAIGN_STATUS_CAMPAIGN_DAILY_BUDGET_EXCEED"
	CampaignStatusOk             = "CAMPAIGN_STATUS_DELIVERY_OK"
	CampaignStatusEnable         = "CAMPAIGN_STATUS_ENABLE"
	CampaignDailyNotExceed       = "CAMPAIGN_DAILY_BUDGET_NOT_EXCEED"
	CampaignDailyExceed          = "CAMPAIGN_DAILY_BUDGET_EXCEED"
	AdGroupStatusDelete          = "ADGROUP_STATUS_DELETE"
	AdGroupStatusDisable         = "ADGROUP_STATUS_DISABLE"
	AdGroupStatusCampaignDisable = "ADGROUP_STATUS_CAMPAIGN_DISABLE"
	AdGroupStatusDone            = "ADGROUP_STATUS_DONE"
	AdGroupStatusAudit           = "ADGROUP_STATUS_AUDIT"
	AdGroupStatusDeny            = "ADGROUP_STATUS_AUDIT_DENY"
	AdGroupStatusNoCreative      = "ADGROUP_STATUS_NO_CREATIVE"
	AdGroupStatusNotStart        = "ADGROUP_STATUS_NOT_START"
	AdGroupStatusFrozen          = "ADGROUP_STATUS_FROZEN"
	AdGroupStatusBalance         = "ADGROUP_STATUS_BALANCE_EXCEED"
	AdGroupStatusBudget          = "ADGROUP_STATUS_ADVERTISER_BUDGET_EXCEED"
	AdGroupStatusCampaignBudget  = "ADGROUP_STATUS_CAMPAIGN_BUDGET_EXCEED"
	AdGroupStatusOk              = "ADGROUP_STATUS_DELIVERY_OK"
	AdGroupStatusEnable          = "ADGROUP_STATUS_ENABLE"
	BalanceNotExceed             = "ADVERTISER_BALANCE_NOT_EXCEED"
	BalanceExceed                = "ADVERTISER_BALANCE_EXCEED"
	FlowResourceShowAd           = "FLOW_RESOURCE_SHOWAD"
	CreativeCategoryThirdParty   = "CREATIVE_SIZE_CATEGORY_THIRD_PARTY"
	CreativeCategorySelf         = "CREATIVE_SIZE_CATEGORY_SELF_OWNED"
	CreativeCategoryOther        = "CREATIVE_SIZE_CATEGORY_OTHER"

	TargetingDictionaryStructLine = 1
	TargetingDictionaryStructTree = 2

	TargetingTypeApp    = "TARGET_TYPE_APP"
	TargetingTypeNotApp = "TARGET_TYPE_NOT_APP"
)

var (
	// OptStatus 计划/任务/创意操作状态  campaign_status
	OptStatus = map[string]string{
		OptStatusEnable:  "投放中",
		OptStatusDisable: "已暂停",
		OptStatusDelete:  "已删除",
	}
	// ProductType 推广产品类型
	ProductType = map[string]string{
		ProductTypeWEB:        "WEB 网页",
		ProductTypeAndroidApp: "应用",
	}
	// CampaignType 计划类型
	CampaignType = map[string]string{
		CampaignTypeDisplay:  "展示广告",
		CampaignTypeShopping: "商品广告",
	}
	// CampaignShowStatus 计划展示状态
	CampaignShowStatus = map[string]string{
		CampaignStatusEnable:      "已启用",
		CampaignStatusOk:          "投放中",
		CampaignStatusDaily:       "未投放(计划到达日预算)",
		CampaignStatusDailyBudget: "未投放(账户到达日预算)",
		CampaignStatusDelete:      "已删除",
		CampaignStatusDisable:     "已暂停",
		CampaignStatusFrozen:      "未投放(账户冻结)",
		CampaignStatusBalance:     "未投放(余额不足)",
	}
	// SyncFlow 同时同步投放搜索广告网络
	SyncFlow = map[string]string{
		SyncFlowYes: "展示广告网络同时投放到搜索广告网络",
		SyncFlowNo:  "只投放展示广告网络",
	}
	// DailyBudgetOpts 计划日限额操作类型
	DailyBudgetOpts = map[string]string{
		DailyBudgetUpdateToday:    "修改当日限额",
		DailyBudgetUpdateTomorrow: "修改计划限额，次日生效",
		DailyBudgetDeleteTomorrow: "删除计划次日日限额",
	}
	// CampaignDaily 计划日预算状态
	CampaignDaily = map[string]string{
		CampaignDailyNotExceed: "未达日预算",
		CampaignDailyExceed:    "到达日预算",
	}
	// AdgroupStatus 任务界面显示的状态
	AdgroupStatus = map[string]string{
		AdGroupStatusDelete:          "已删除",
		AdGroupStatusDisable:         "已暂停",
		AdGroupStatusCampaignDisable: "计划暂停",
		AdGroupStatusDone:            "投放结束",
		AdGroupStatusAudit:           "审核中",
		AdGroupStatusDeny:            "审核不通过",
		AdGroupStatusNoCreative:      "待上传创意",
		AdGroupStatusNotStart:        "未投放(未到投放日期)",
		AdGroupStatusFrozen:          "未投放(已冻结)",
		AdGroupStatusBalance:         "未投放(余额不足)",
		AdGroupStatusBudget:          "未投放(到达日预算)",
		AdGroupStatusCampaignBudget:  "未投放(达到日预算)",
		AdGroupStatusOk:              "投放中",
		AdGroupStatusEnable:          "已启用",
	}
	// BalanceStatus 账户余额状态
	BalanceStatus = map[string]string{
		BalanceNotExceed: "余额充足",
		BalanceExceed:    "余额不足",
	}
	// FlowResource 投放网络
	FlowResource = map[string]string{
		FlowResourceShowAd: "展示广告网络",
	}
	// CreativeCategory 任务版位分类
	CreativeCategory = map[string]string{
		CreativeCategoryThirdParty: "三方媒体资源",
		CreativeCategorySelf:       "自有媒体资源",
		CreativeCategoryOther:      "其他首选资",
	}
	// TargetingDictionaryKeys 字典 key 列表
	TargetingDictionaryKeys = []string{
		"app_category",
		"series_type",
		"carrier",
		"location_category",
		"pre_define_audience",
		"not_pre_define_audience",
		"gender",
		"age",
		"device_price",
		"network_type",
		"pre_define_audience",
		"not_pre_define_audience",
		"media_app_category",
		"language",
		"app_interest",
	}
	TargetingDictionaryStruct = map[int64][]string{
		TargetingDictionaryStructLine: {
			"pre_define_audience",
			"not_pre_define_audience",
			"gender",
			"age",
			"device_price",
			"network_type",
			"pre_define_audience",
			"not_pre_define_audience",
			"media_app_category",
			"language",
			"app_interest",
			"location_category",
		},
		TargetingDictionaryStructTree: {"carrier", "app_category", "series_type"},
	}
	// TargetingType 定向包类型
	TargetingType = map[string]string{
		TargetingTypeApp:    "应用类",
		TargetingTypeNotApp: "非应用类",
	}
)
