package statements

type TargetingValue struct {
	Value []string `json:"value"`
}

type TargetingCreateResp struct {
	BaseAdsResp
	Data struct {
		TargetingId int64 `json:"targeting_id"`
	} `json:"data"`
}

type TrackingRequest struct {
	Filtering    TrackingFilter `json:"filtering"`
	AdvertiserId string         `json:"advertiser_id"`
	Page         int            `json:"page"`
	PageSize     int            `json:"page_size"`
}
type TrackingFilter struct {
	ProductUniqueFlag string `json:"product_unique_flag"`
}
type TrackingResponse struct {
	BaseAdsResp
	Data TrackingData `json:"data"`
}
type TrackingData struct {
	Total int64       `json:"total"`
	Data  []*Tracking `json:"data"`
}
type Tracking struct {
	ClickTrackingUrl string `json:"click_tracking_url"`
	TrackingStatus   string `json:"tracking_status"`
	EffectName       string `json:"effect_name"`
	ImpTrackingUrl   string `json:"imp_tracking_url"`
	EffectType       string `json:"effect_type"`
	TrackingId       int64  `json:"tracking_id"`
}
