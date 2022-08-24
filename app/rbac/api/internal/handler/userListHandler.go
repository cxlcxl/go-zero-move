package handler

import (
	"net/http"

	"business/app/rbac/api/internal/logic"
	"business/app/rbac/api/internal/svc"
	"business/app/rbac/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func userListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserListLogic(r.Context(), svcCtx)
		resp, err := l.UserList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
