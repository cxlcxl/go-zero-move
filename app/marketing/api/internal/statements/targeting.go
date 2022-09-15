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
