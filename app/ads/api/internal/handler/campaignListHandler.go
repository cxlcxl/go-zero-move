package handler

import (
	"net/http"

	"business/app/ads/api/internal/logic"
	"business/app/ads/api/internal/svc"
	"business/app/ads/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func campaignListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CampaignListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCampaignListLogic(r.Context(), svcCtx)
		resp, err := l.CampaignList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
