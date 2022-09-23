package handler

import (
	"net/http"

	"business/app/marketing/api/internal/logic"
	"business/app/marketing/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func assetDimensionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewAssetDimensionLogic(r.Context(), svcCtx)
		resp, err := l.AssetDimension()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
