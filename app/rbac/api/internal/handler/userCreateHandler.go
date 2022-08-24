package handler

import (
	"net/http"

	"business/app/rbac/api/internal/logic"
	"business/app/rbac/api/internal/svc"
	"business/app/rbac/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func userCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserCreateLogic(r.Context(), svcCtx)
		resp, err := l.UserCreate(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
