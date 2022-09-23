package handler

import (
	"business/app/marketing/api/internal/logic"
	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func assetUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AssetUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAssetUploadLogic(r, svcCtx)
		resp, err := l.AssetUpload(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
