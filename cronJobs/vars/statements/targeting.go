package statements

type TargetingRequest struct {
	RequestPage
	AdvertiserId string `json:"advertiser_id"`
}

type TargetingValue struct {
	Value []string `json:"value"`
}

type Targeting struct {
	TargetingId           int64          `json:"targeting_id"`
	TargetingType         string         `json:"targeting_type"`
	TargetingName         string         `json:"targeting_name"`
	Gender                TargetingValue `json:"gender_struct"`
	Age                   TargetingValue `json:"age_struct"`
	InstalledApps         TargetingValue `json:"installed_apps_struct"`
	AppCategoryInstalled  TargetingValue `json:"app_category_installed_struct"`
	NetworkType           TargetingValue `json:"network_type_struct"`
	Audience              TargetingValue `json:"audience_struct"`
	NotAudience           TargetingValue `json:"not_audience_struct"`
	SeriesType            TargetingValue `json:"series_type_struct"`
	AppCategoryOfMedia    TargetingValue `json:"app_category_of_media_struct"`
	Carrier               TargetingValue `json:"carrier_struct"`
	Language              TargetingValue `json:"language_struct"`
	HighAppInterest       TargetingValue `json:"high_app_interest_struct"`
	CurrentCustomLocation TargetingValue `json:"current_custom_location_struct"`
}
type TargetingResponse struct {
	BaseResponse
	Data struct {
		Total      int64        `json:"total"`
		Targetings []*Targeting `json:"data"`
	} `json:"data"`
}
