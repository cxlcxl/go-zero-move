package logic

import (
	"business/common/curl"
	"business/cronJobs/jobs/country/model"
	"encoding/json"
	"time"
)

type AccessToken struct {
	Error            int64  `json:"error"`
	ErrorDescription string `json:"error_description"`
	AccessToken      string `json:"access_token"`
	ExpiresIn        int64  `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
}

// GetHWAdsPostResponse 调用华为 POST 类接口
func GetHWAdsPostResponse(url string, data, res interface{}, token string) (err error) {
	request, err := curl.New(url).Post().JsonData(data)
	if err != nil {
		return
	}
	err = request.Request(res, curl.Authorization(token))
	if err != nil {
		return
	}
	return
}

// GetAdsModel 获取 ads 写入日志模型
func GetAdsModel(statDate, apiModule string, accountId, page, totalPage, totalNum, pageSize int64, adsLog *model.AdsRequestLogs, data interface{}) *model.AdsRequestLogs {
	var isCompleted int64 = 0
	nextPage := page
	if totalPage <= page {
		isCompleted = 1
	} else {
		nextPage += 1
	}
	var id int64 = 0
	if adsLog != nil {
		id = adsLog.Id
	}
	marshal, _ := json.Marshal(data)
	d, _ := time.Parse("2006-01-02", statDate)
	adsModel := &model.AdsRequestLogs{
		Id:                 id,
		StatDay:            d,
		ApiModule:          apiModule,
		AccountId:          accountId,
		RequestJsonBody:    string(marshal),
		CurrentRequestPage: page,
		NextRequestPage:    nextPage,
		IsCompleted:        isCompleted,
		TotalPage:          totalPage,
		TotalNum:           totalNum,
		PageSize:           pageSize,
		LastRequestTime:    time.Now(),
	}
	return adsModel
}
