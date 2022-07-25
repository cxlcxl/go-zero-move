// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"business/app/rbac/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: loginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/profile",
					Handler: profileHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware, serverCtx.PermissionMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/user/list",
					Handler: userListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/create",
					Handler: userCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/update",
					Handler: userUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/info",
					Handler: getUserInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/destroy",
					Handler: UserDestroyHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/role/list",
					Handler: roleListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role/create",
					Handler: roleCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role/update",
					Handler: roleUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/role/info",
					Handler: getRoleInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role/destroy",
					Handler: RoleDestroyHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
