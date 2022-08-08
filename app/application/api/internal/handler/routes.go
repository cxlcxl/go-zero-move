// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"business/app/application/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware, serverCtx.PermissionMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: appCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: appUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/info",
					Handler: getAppInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: getAppListHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/app"),
	)
}