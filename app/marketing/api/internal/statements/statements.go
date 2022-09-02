package statements

type BaseAdsResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type CampaignCreate struct {
	AdvertiserId string `json:"advertiser_id"`
	CampaignName string `json:"campaign_name"`
	ProductType  string `json:"product_type"`
	DailyBudget  int64  `json:"daily_budget"`
	CampaignType string `json:"campaign_type"`
	SyncFlow     string `json:"sync_flow_resource_searchad"`
}

type CampaignCreateResp struct {
	*BaseAdsResp
	Data struct {
		CampaignId string `json:"campaign_id"`
	} `json:"data"`
}
