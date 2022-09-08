package statements

type DictionaryRequest struct {
	Language      string   `json:"language"`
	TargetingList []string `json:"targeting_list"`
}

type DictionaryResponse struct {
	BaseResponse
	Data struct {
		TreeTargetingMap struct {
			Carrier []struct {
				Id    string `json:"id"`
				Pid   string `json:"pid"`
				Label string `json:"label"`
				Value string `json:"value"`
				Code  string `json:"code"`
			} `json:"carrier"`
			AppCategory []struct {
				Id    string `json:"id"`
				Pid   string `json:"pid"`
				Label string `json:"label"`
				Value string `json:"value"`
			} `json:"app_category"`
			SeriesType []struct {
				Id    string `json:"id"`
				Pid   string `json:"pid"`
				Label string `json:"label"`
				Value string `json:"value"`
			} `json:"series_type"`
		} `json:"tree_targeting_map"`
		LinearTargetingMap struct {
			DevicePrice          []interface{} `json:"device_price"`
			NotPreDefineAudience []struct {
				Label string `json:"label"`
				Value string `json:"value"`
				Id    string `json:"id"`
			} `json:"not_pre_define_audience"`
			Gender []struct {
				Label string `json:"label"`
				Value string `json:"value"`
			} `json:"gender"`
			PreDefineAudience []struct {
				Label string `json:"label"`
				Value string `json:"value"`
				Id    string `json:"id"`
			} `json:"pre_define_audience"`
			Language []struct {
				Label string `json:"label"`
				Value string `json:"value"`
			} `json:"language"`
			MediaAppCategory []struct {
				Label string `json:"label"`
				Value string `json:"value"`
				Id    string `json:"id"`
				Pid   string `json:"pid"`
			} `json:"media_app_category"`
			AppInterest []struct {
				Label string `json:"label"`
				Value string `json:"value"`
				Id    string `json:"id"`
				Seq   string `json:"seq"`
				Pid   string `json:"pid,omitempty"`
			} `json:"app_interest"`
			NetworkType []struct {
				Label string `json:"label"`
				Value string `json:"value"`
			} `json:"network_type"`
			Age []struct {
				Label string `json:"label"`
				Value string `json:"value"`
			} `json:"age"`
		} `json:"linear_targeting_map"`
	} `json:"data"`
}
