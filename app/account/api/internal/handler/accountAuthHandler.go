package handler

import (
	"net/http"

	"business/app/account/api/internal/logic"
	"business/app/account/api/internal/svc"
	"business/app/account/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func accountAuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AccountInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAccountAuthLogic(r.Context(), svcCtx)
		resp, err := l.AccountAuth(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
