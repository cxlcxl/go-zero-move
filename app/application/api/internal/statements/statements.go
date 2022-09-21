package statements

type AppPullRequest struct {
	AdvertiserId string       `json:"advertiser_id"`
	Page         int          `json:"page"`
	PageSize     int          `json:"page_size"`
	Filtering    ReqFiltering `json:"filtering"`
}
type ReqFiltering struct {
	ProductType string `json:"product_type"`
	AgAppType   string `json:"ag_app_type"`
}
type ProductInfo struct {
	ProductType string  `json:"product_type"`
	ProductId   string  `json:"product_id"`
	ProductInfo AppInfo `json:"product_info"`
}
type AppInfo struct {
	App App `json:"app"`
}
type App struct {
	IconUrl             string `json:"icon_url"`
	AppStoreDownloadUrl string `json:"app_store_download_url"`
	PackageName         string `json:"package_name"`
	Description         string `json:"description"`
	AppId               string `json:"app_id"`
	ProductName         string `json:"product_name"`
}

type AppAdsResponse struct {
	Code    string     `json:"code"`
	Message string     `json:"message"`
	Data    AdsAppInfo `json:"data"`
}

type AdsAppInfo struct {
	Total int64          `json:"total"`
	Data  []*ProductInfo `json:"data"`
}
