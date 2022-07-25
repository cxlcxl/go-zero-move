package handler

import (
	"business/app/rbac/api/internal/logic"
	"business/app/rbac/api/internal/svc"
	"business/app/rbac/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func userUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserUpdateLogic(r.Context(), svcCtx)
		resp, err := l.UserUpdate(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
