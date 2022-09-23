package handler

import (
	"net/http"

	"business/app/marketing/api/internal/logic"
	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func assetListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AssetListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAssetListLogic(r.Context(), svcCtx)
		resp, err := l.AssetList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
