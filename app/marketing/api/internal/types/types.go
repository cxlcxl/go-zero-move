// Code generated by goctl. DO NOT EDIT.
package types

type BaseResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type PromotionCreateReq struct {
	CampaignName  string `json:"campaign_name" validate:"required"` // 计划名称
	ProductType   string `json:"product_type" validate:"required"`  // 推广产品
	DailyBudget   int64  `json:"daily_budget" validate:"numeric"`   // 日预算
	AdvertiserId  string `json:"advertiser_id"`                     // 广告主ID
	MarketingGoal string `json:"marketing_goal"`                    // 营销目标，默认为无目的
}

type PromotionCreateInfo struct {
	CampaignId string `json:"campaign_id"`
}

type PromotionCreateResp struct {
	Code int64                `json:"code"`
	Msg  string               `json:"msg"`
	Data *PromotionCreateInfo `json:"data"`
}

type PromotionListReq struct {
	Page         int64  `form:"page"`
	PageSize     int64  `form:"page_size"`
	CampaignName string `form:"campaign_name,optional"`
	CampaignId   string `form:"campaign_id,optional"`
	ShowStatus   string `form:"show_status,optional"`
}

type Promotion struct {
	Id                        int64  `json:"id"`
	CampaignName              string `json:"campaign_name"`                // 计划名称
	CampaignType              string `json:"campaign_type"`                // 计划类型
	FlowResource              string `json:"flow_resource"`                // 投放网络
	ProductType               string `json:"product_type"`                 // 推广产品
	UserBalanceStatus         string `json:"user_balance_status"`          // 账户余额状态
	CampaignStatus            string `json:"campaign_status"`              // 操作状态
	CampaignDailyBudgetStatus string `json:"campaign_daily_budget_status"` // 计划日预算状态
	ShowStatus                string `json:"show_status"`                  // 计划状态
	CreatedAt                 int64  `json:"created_at"`                   // 计划创建的时间
	MarketingGoal             string `json:"marketing_goal"`               //
	CampaignId                string `json:"campaign_id"`                  // 计划ID
	TodayDailyBudget          int64  `json:"today_daily_budget"`           // 当日计划日限额
	TomorrowDailyBudget       int64  `json:"tomorrow_daily_budget"`        // 次日计划日限额，不返回表示与当日计划日限额相同
	SyncFlowResourceSearchad  string `json:"sync_flow_resource_searchad"`  // 同时同步投放搜索广告网络 YES|NO
}

type PromotionListResp struct {
	BaseResp
	Total int64        `json:"total"`
	Data  []*Promotion `json:"data"`
}
