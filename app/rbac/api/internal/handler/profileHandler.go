package handler

import (
	"net/http"

	"business/app/rbac/api/internal/logic"
	"business/app/rbac/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func profileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewProfileLogic(r.Context(), svcCtx)
		resp, err := l.Profile()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
