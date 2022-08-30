package vars

const (
	OptStatusEnable  = "OPERATION_STATUS_ENABLE"
	OptStatusDisable = "OPERATION_STATUS_DISABLE"
	OptStatusDelete  = "OPERATION_STATUS_DELETE"

	ProductTypeWEB        = "WEB"
	ProductTypeAndroidApp = "ANDROID_APP"

	CampaignTypeDisplay  = "CAMPAIGN_TYPE_DISPLAY"
	CampaignTypeShopping = "CAMPAIGN_TYPE_SHOPPING"

	SyncFlowYes = "YES"
	SyncFlowNo  = "NO"
)

var (
	// OptStatus 计划/任务/创意操作状态
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
	// SyncFlow 同时同步投放搜索广告网络
	SyncFlow = map[string]string{
		SyncFlowYes: "展示广告网络同时投放到搜索广告网络",
		SyncFlowNo:  "只投放展示广告网络",
	}
)
