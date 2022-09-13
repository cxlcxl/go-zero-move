// Code generated by goctl. DO NOT EDIT.
package types

type Dictionary struct {
	DictKey  string        `json:"dict_key"`
	Id       string        `json:"id"`
	Pid      string        `json:"pid"`
	Label    string        `json:"label"`
	Value    string        `json:"value"`
	Code     string        `json:"code"`
	Seq      string        `json:"seq"`
	Children []*Dictionary `json:"children"`
}

type DictReq struct {
	DictKey string `form:"dict_key"`
}

type DictResp struct {
	BaseResp
	Data map[string][]*Dictionary `json:"data"`
}

type BaseResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type CampaignCreateReq struct {
	AccountId    int64  `json:"account_id"`
	CampaignName string `json:"campaign_name" validate:"required"`               // 计划名称
	ProductType  string `json:"product_type" validate:"required,product_type"`   // 推广产品
	DailyBudget  int64  `json:"daily_budget" validate:"numeric"`                 // 日预算
	CampaignType string `json:"campaign_type" validate:"required,campaign_type"` // 计划类型
	SyncFlow     string `json:"sync_flow_resource_searchad"`                     // 同时同步投放搜索广告网络
}

type CampaignCreateInfo struct {
	CampaignId string `json:"campaign_id"`
}

type CampaignCreateResp struct {
	Code int64               `json:"code"`
	Msg  string              `json:"msg"`
	Data *CampaignCreateInfo `json:"data"`
}

type CampaignListReq struct {
	Page         int64  `form:"page"`
	PageSize     int64  `form:"page_size"`
	CampaignName string `form:"campaign_name,optional"`
	CampaignId   string `form:"campaign_id,optional"`
	CampaignType string `form:"show_status,optional"`
}

type Campaign struct {
	Id                        int64  `json:"id"`
	CampaignName              string `json:"campaign_name"`                // 计划名称
	CampaignType              string `json:"campaign_type"`                // 计划类型
	FlowResource              string `json:"flow_resource"`                // 投放网络
	ProductType               string `json:"product_type"`                 // 推广产品
	UserBalanceStatus         string `json:"user_balance_status"`          // 账户余额状态
	OptStatus                 string `json:"opt_status"`                   // 操作状态
	CampaignDailyBudgetStatus string `json:"campaign_daily_budget_status"` // 计划日预算状态
	ShowStatus                string `json:"show_status"`                  // 计划状态
	MarketingGoal             string `json:"marketing_goal"`               //
	CampaignId                string `json:"campaign_id"`                  // 计划ID
	TodayDailyBudget          int64  `json:"today_daily_budget"`           // 当日计划日限额
	TomorrowDailyBudget       int64  `json:"tomorrow_daily_budget"`        // 次日计划日限额，不返回表示与当日计划日限额相同
	SyncFlowResourceSearchad  string `json:"sync_flow_resource_searchad"`  // 同时同步投放搜索广告网络 YES|NO
	IsCallback                int64  `json:"is_callback"`
	CreatedAt                 int64  `json:"created_at"` // 计划创建的时间
	UpdatedAt                 int64  `json:"updated_at"`
}

type CampaignListResp struct {
	BaseResp
	Data CampaignListData `json:"data"`
}

type CampaignListData struct {
	Campaigns []*Campaign `json:"campaigns"`
	Total     int64       `json:"total"`
}

type Resources struct {
	OptStatus       map[string]string `json:"opt_status"`
	ProductType     map[string]string `json:"product_type"`
	CampaignType    map[string]string `json:"campaign_type"`
	SyncFlow        map[string]string `json:"sync_flow"`
	DailyBudgetOpts map[string]string `json:"daily_budget_opts"`
	ShowStatus      map[string]string `json:"show_status"`
	CampaignDaily   map[string]string `json:"campaign_daily"`
	BalanceStatus   map[string]string `json:"balance_status"`
	FlowResource    map[string]string `json:"flow_resource"`
}

type CampaignResources struct {
	BaseResp
	Data *Resources `json:"data"`
}

type TargetingCreateReq struct {
	AccountId                int64    `json:"account_id"`
	TargetingType            string   `json:"targeting_type"`
	TargetingName            string   `json:"targeting_name"`
	GenderStruct             string   `json:"gender_struct"`
	AgeStruct                []string `json:"age_struct"`
	NetworkTypeStruct        []string `json:"network_type_struct"`
	AppBehavior              string   `json:"app_behavior"`
	AppBehaviors             []string `json:"app_behaviors"`
	AppInterest              string   `json:"app_interest"`
	AppInterests             []string `json:"app_interests"`
	Audience                 string   `json:"audience"`
	AudienceStruct           []string `json:"audience_struct"`
	NotAudienceStruct        []string `json:"not_audience_struct"`
	SeriesType               string   `json:"series_type"`
	Series                   []string `json:"series"`
	MediaAppCategory         string   `json:"media_app_category"`
	AppCategoryOfMediaStruct []string `json:"app_category_of_media_struct"`
	LanguageCheck            string   `json:"language_check"`
	Language                 []string `json:"language"`
	InstalledApps            string   `json:"installed_apps"`
}

type TargetingCreateRsInfo struct {
	TargetingId string `json:"targeting_id"`
}

type TargetingCreateResp struct {
	BaseResp
	Data TargetingCreateRsInfo `json:"data"`
}
