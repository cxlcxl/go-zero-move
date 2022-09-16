package statements

const (
	ClientId = "104661969"
)

type BaseResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Pages struct {
	Page      int64 `json:"page"`
	PageSize  int64 `json:"page_size"`
	TotalNum  int64 `json:"total_num"`
	TotalPage int64 `json:"total_page"`
}

type RequestPage struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"page_size"`
}
