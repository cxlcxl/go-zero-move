package handler

import (
	"net/http"

	"business/app/account/api/internal/logic"
	"business/app/account/api/internal/svc"
	"business/app/account/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func accessTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AccessTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAccessTokenLogic(r.Context(), svcCtx)
		resp, err := l.AccessToken(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
