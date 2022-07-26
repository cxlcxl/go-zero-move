// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"business/app/account/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware, serverCtx.PermissionMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/token",
					Handler: accessTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/refresh",
					Handler: refreshTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: accountListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/auth",
					Handler: accountAuthHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/info",
					Handler: accountInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: accountCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: accountUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/search",
					Handler: accountSearchHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/default",
					Handler: accountDefaultListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/parents",
					Handler: accountDefaultParentsHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/account"),
	)
}
