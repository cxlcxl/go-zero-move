// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"business/app/marketing/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware, serverCtx.PermissionMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/campaign/resources",
					Handler: campaignResourcesHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/campaign/create",
					Handler: campaignCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/campaign/list",
					Handler: campaignListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/dictionary/query",
					Handler: dictionaryQueryHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/marketing"),
	)
}
