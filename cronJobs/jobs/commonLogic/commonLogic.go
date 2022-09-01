package logic

import (
	"business/common/curl"
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
