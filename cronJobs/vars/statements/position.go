package statements

type PositionRequest struct {
	AdvertiserId string            `json:"advertiser_id"`
	Filtering    PositionFiltering `json:"filtering"`
}
type PositionFiltering struct {
	Category string `json:"category"`
}

type PositionResponse struct {
	BaseResponse
	Data struct {
		CreativeSizeInfoList []*CreativeSizeInfo `json:"creative_size_info_list"`
	} `json:"data"`
}

type CreativeSizeInfo struct {
	CreativeSizeId       int `json:"creative_size_id"`
	CreativeSizeBaseInfo struct {
		CreativeSizeNameDsp       string                   `json:"creative_size_name_dsp"`
		CreativeSizeDescription   string                   `json:"creative_size_description"`
		CreativeSizeSampleList    []*CreativeSizeSample    `json:"creative_size_sample_list"`    // 预览图
		CreativeSizePlacementList []*CreativeSizePlacement `json:"creative_size_placement_list"` // 版位规格
	} `json:"creative_size_base_info"`
	CreativeSizeOperationInfo struct {
		SupportProductType         string `json:"support_product_type"`
		SupportObjectiveType       string `json:"support_objective_type"`
		IsSupportTimePeriod        string `json:"is_support_time_period"`
		IsSupportMultipleCreatives string `json:"is_support_multiple_creatives"`
	} `json:"creative_size_operation_info"`
	CreativeSizePriceInfo struct {
		SupportPriceType string `json:"support_price_type"`
	} `json:"creative_size_price_info"`
}

type CreativeSizeSample struct {
	CreativeSizeSample string `json:"creative_size_sample"`
	PreviewTitle       string `json:"preview_title"`
}
type CreativeSizePlacement struct {
	PlacementSizeId            string `json:"placement_size_id"`
	CreativeSize               string `json:"creative_size"`
	CreativeSizeSubType        string `json:"creative_size_sub_type"`
	IsSupportMultipleCreatives string `json:"is_support_multiple_creatives"`
}
