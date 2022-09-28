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

type PositionElementRequest struct {
	AdvertiserId string                `json:"advertiser_id"`
	Filtering    PositionElementFilter `json:"filtering"`
}
type PositionElementFilter struct {
	CreativeSizeId string `json:"creative_size_id"`
}

type PositionElementResponse struct {
	BaseResponse
	Data PositionElement `json:"data"`
}

type PositionElement struct {
	CreativeSizeId  int                `json:"creative_size_id"`
	ElementInfoList []*ElementInfoList `json:"placement_size_elementinfolist"`
}

type ElementInfoList struct {
	Subtype         string         `json:"creative_size_subtype"`
	ElementInfoList []*ElementInfo `json:"creative_element_info_list"`
}

type ElementInfo struct {
	ElementId       int               `json:"creative_size_element_id"`
	ElementName     string            `json:"creative_size_element_name"`
	ElementTitle    string            `json:"creative_size_element_title"`
	ElementCaption  string            `json:"creative_size_element_caption"`
	Width           int64             `json:"width"`
	Height          int64             `json:"height"`
	MinWidth        int64             `json:"min_width"`
	MinHeight       int64             `json:"min_height"`
	MinLength       int64             `json:"min_length"`
	MaxLength       int64             `json:"max_length"`
	FileSizeKbLimit int64             `json:"file_size_kb_limit"`
	GifSizeKbLimit  int64             `json:"gif_size_kb_limit"`
	FileFormat      string            `json:"file_format"`
	Pattern         string            `json:"pattern"`
	Duration        []ElementDuration `json:"duration"`
	MinOccurs       string            `json:"min_occurs"`
	MaxOccurs       string            `json:"max_occurs"`
}
type ElementDuration struct {
	Min int64 `json:"min"`
	Max int64 `json:"max"`
}

type CreativeSizePriceRequest struct {
	AdvertiserId string                  `json:"advertiser_id"`
	Filtering    CreativeSizePriceFilter `json:"filtering"`
}
type CreativeSizePriceFilter struct {
	CreativeSizeId string `json:"creative_size_id"`
	PriceType      string `json:"price_type"`
}

type CreativeSizePriceResponse struct {
	BaseResponse
	Data CreativeSizePrice `json:"data"`
}
type CreativeSizePrice struct {
	FloorPrice float64 `json:"floor_price"`
}
