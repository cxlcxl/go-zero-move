package handler

import (
	"net/http"

	"business/app/application/api/internal/logic"
	"business/app/application/api/internal/svc"
	"business/app/application/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func appPullHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AppPullReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAppPullLogic(r.Context(), svcCtx)
		resp, err := l.AppPull(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
