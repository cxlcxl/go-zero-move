package statements

type PositionElementRequest struct {
	AdvertiserId string                `json:"advertiser_id"`
	Filtering    PositionElementFilter `json:"filtering"`
}
type PositionElementFilter struct {
	CreativeSizeId string `json:"creative_size_id"`
}
type PositionElementResponse struct {
	BaseAdsResp
	Data PositionElement `json:"data"`
}

type PositionElement struct {
	CreativeSizeId  int64              `json:"creative_size_id"`
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
	Width           int               `json:"width"`
	Height          int               `json:"height"`
	FileSizeKbLimit int               `json:"file_size_kb_limit"`
	FileFormat      string            `json:"file_format"`
	Duration        []ElementDuration `json:"duration"`
	MinOccurs       string            `json:"min_occurs"`
	MaxOccurs       string            `json:"max_occurs"`
	SelectTexts     []ElementLanguage `json:"creative_element_select_texts"`
}
type ElementDuration struct {
	Min int64 `json:"min"`
	Max int64 `json:"max"`
}
type ElementLanguage struct {
	Language string `json:"language"`
}
