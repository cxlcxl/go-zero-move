package statements

type AdsAssetRequest struct {
	AdvertiserId string    `json:"advertiser_id"`
	Page         int       `json:"page"`
	PageSize     int64     `json:"page_size"`
	Filtering    Filtering `json:"filtering"`
}

type Filtering struct {
	AssetStatus string `json:"asset_status"`
}

type AdsAssetResponse struct {
	BaseAdsResp
	Data AdsAssetInfo `json:"data"`
}

type AdsAssetInfo struct {
	Total              int64        `json:"total"`
	CreativeAssetInfos []*AssetInfo `json:"creative_asset_infos"`
}

type AssetInfo struct {
	AssetStatus       string `json:"asset_status"`
	AssetId           int64  `json:"asset_id"`
	AssetName         string `json:"asset_name"`
	AssetType         string `json:"asset_type"`
	FileUrl           string `json:"file_url"`
	Width             int64  `json:"width"`
	Height            int64  `json:"height"`
	FileSize          int64  `json:"file_size"`
	FileFormat        string `json:"file_format"`
	FileHashSha256    string `json:"file_hash_sha256"`
	VideoPlayDuration int64  `json:"video_play_duration"`
}

type AdsAssetDeleteResponse struct {
	BaseAdsResp
	Data []AssetDeleteRsSlice `json:"data"`
}
type AssetDeleteRsSlice struct {
	ErrorCode string `json:"error_code"`
	AssetId   int64  `json:"asset_id"`
	Message   string `json:"message"`
}

type FileTokenReq struct {
	AdvertiserId string `json:"advertiser_id"`
}
type FileTokenResponse struct {
	BaseAdsResp
	Data struct {
		FileToken string `json:"file_token"`
	} `json:"data"`
}

type AssetCreateResponse struct {
	BaseAdsResp
	Data struct {
		Url     string `json:"url"`
		AssetId int64  `json:"asset_id"`
	} `json:"data"`
}

type AssetDimension struct {
	Width  int64
	Height int64
}

var (
	PictureDimensions = []AssetDimension{
		{Width: 120, Height: 120},
		{Width: 160, Height: 160},
		{Width: 192, Height: 192},
		{Width: 210, Height: 150},
		{Width: 216, Height: 216},
		{Width: 225, Height: 150},
		{Width: 250, Height: 250},
		{Width: 256, Height: 256},
		{Width: 285, Height: 120},
		{Width: 294, Height: 222},
		{Width: 300, Height: 50},
		{Width: 300, Height: 250},
		{Width: 312, Height: 432},
		{Width: 320, Height: 50},
		{Width: 320, Height: 480},
		{Width: 334, Height: 136},
		{Width: 336, Height: 280},
		{Width: 392, Height: 392},
		{Width: 400, Height: 400},
		{Width: 450, Height: 450},
		{Width: 480, Height: 320},
		{Width: 480, Height: 800},
		{Width: 486, Height: 216},
		{Width: 486, Height: 468},
		{Width: 512, Height: 512},
		{Width: 576, Height: 720},
		{Width: 608, Height: 608},
		{Width: 612, Height: 468},
		{Width: 640, Height: 184},
		{Width: 640, Height: 256},
		{Width: 640, Height: 360},
		{Width: 640, Height: 960},
		{Width: 720, Height: 1280},
		{Width: 720, Height: 1440},
		{Width: 728, Height: 90},
		{Width: 768, Height: 1024},
		{Width: 880, Height: 552},
		{Width: 900, Height: 750},
		{Width: 960, Height: 150},
		{Width: 960, Height: 300},
		{Width: 984, Height: 422},
		{Width: 1024, Height: 512},
		{Width: 1024, Height: 768},
		{Width: 1045, Height: 224},
		{Width: 1045, Height: 396},
		{Width: 1080, Height: 96},
		{Width: 1080, Height: 150},
		{Width: 1080, Height: 170},
		{Width: 1080, Height: 270},
		{Width: 1080, Height: 310},
		{Width: 1080, Height: 430},
		{Width: 1080, Height: 432},
		{Width: 1080, Height: 504},
		{Width: 1080, Height: 607},
		{Width: 1080, Height: 1620},
		{Width: 1080, Height: 1920},
		{Width: 1200, Height: 420},
		{Width: 1200, Height: 627},
		{Width: 1248, Height: 400},
		{Width: 1248, Height: 534},
		{Width: 1280, Height: 720},
		{Width: 1280, Height: 1665},
		{Width: 1312, Height: 768},
		{Width: 1334, Height: 222},
		{Width: 1334, Height: 288},
		{Width: 1334, Height: 539},
		{Width: 1404, Height: 180},
		{Width: 1440, Height: 2560},
		{Width: 1920, Height: 1080},
		{Width: 2022, Height: 486},
		{Width: 2184, Height: 270},
		{Width: 2200, Height: 3120},
		{Width: 2496, Height: 468},
		{Width: 2934, Height: 3306},
	}
	VideoDimensions = []AssetDimension{
		{Width: 320, Height: 480},
		{Width: 480, Height: 320},
		{Width: 480, Height: 640},
		{Width: 640, Height: 360},
		{Width: 640, Height: 480},
		{Width: 720, Height: 1080},
		{Width: 720, Height: 1280},
		{Width: 720, Height: 1440},
		{Width: 880, Height: 552},
		{Width: 1080, Height: 1920},
		{Width: 1280, Height: 720},
		{Width: 1920, Height: 1080},
		{Width: 2934, Height: 3306},
	}
)
