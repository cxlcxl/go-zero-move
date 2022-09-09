package statements

type DictionaryRequest struct {
	Language      string   `json:"language"`
	TargetingList []string `json:"targeting_list"`
}

type DictionaryItem struct {
	Id    string `json:"id"`
	Pid   string `json:"pid"`
	Label string `json:"label"`
	Value string `json:"value"`
	Code  string `json:"code"`
	Seq   string `json:"seq"`
}

type DictionaryResponse struct {
	BaseResponse
	Data struct {
		TreeTargetingMap struct {
			Carrier     []*DictionaryItem `json:"carrier"`
			AppCategory []*DictionaryItem `json:"app_category"`
			SeriesType  []*DictionaryItem `json:"series_type"`
		} `json:"tree_targeting_map"`
		LinearTargetingMap struct {
			DevicePrice          []*DictionaryItem `json:"device_price"`
			NotPreDefineAudience []*DictionaryItem `json:"not_pre_define_audience"`
			Gender               []*DictionaryItem `json:"gender"`
			PreDefineAudience    []*DictionaryItem `json:"pre_define_audience"`
			Language             []*DictionaryItem `json:"language"`
			MediaAppCategory     []*DictionaryItem `json:"media_app_category"`
			AppInterest          []*DictionaryItem `json:"app_interest"`
			NetworkType          []*DictionaryItem `json:"network_type"`
			Age                  []*DictionaryItem `json:"age"`
		} `json:"linear_targeting_map"`
	} `json:"data"`
}
