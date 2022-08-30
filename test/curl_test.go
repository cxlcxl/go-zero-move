package test

import (
	"business/common/curl"
	"testing"
)

func TestCurl(t *testing.T) {
	var aa interface{}
	req, err := curl.New("https://ads-dra.cloud.huawei.com/ads/v1/promotion/campaign/create").Post().JsonData(map[string]interface{}{
		"advertiser_id":               "494043463919075840",
		"campaign_name":               "测试测试-阿萨德法撒旦法",
		"product_type":                "WEB",
		"daily_budget":                10,
		"sync_flow_resource_searchad": "YES",
		"campaign_type":               "CAMPAIGN_TYPE_SHOPPING",
	})
	if err != nil {
		t.Fatal(err)
	}
	err = req.Request(&aa, curl.Authorization("Bearer DAECANpUaFqRDp7R4F11i706oEeUo3rG5khnRW9z16BCNXYMQgyOGd/U0wf6DGkzOrGD2ufo28+eLjT07+KNrtX1Xg+aFblhGyaMfdCBG8pB0uaEORWLrKl4RlQTbCZW+g=="), curl.JsonHeader())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(aa)
}
