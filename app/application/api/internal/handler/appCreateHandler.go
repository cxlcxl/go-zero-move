package handler

import (
	"net/http"

	"business/app/application/api/internal/logic"
	"business/app/application/api/internal/svc"
	"business/app/application/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func appCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AppCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAppCreateLogic(r.Context(), svcCtx)
		resp, err := l.AppCreate(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
