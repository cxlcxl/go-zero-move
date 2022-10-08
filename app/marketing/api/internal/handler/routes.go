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
					Path:    "/campaign/info",
					Handler: campaignInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/campaign/list",
					Handler: campaignListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/campaign/bind-app",
					Handler: campaignBindAppHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/tracking/list",
					Handler: trackingListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/tracking/refresh",
					Handler: trackingRefreshHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/dictionary/query",
					Handler: dictionaryQueryHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/targeting/location",
					Handler: targetingLocationHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/targeting/list",
					Handler: targetingListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/targeting/create",
					Handler: targetingCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/resource/pricing",
					Handler: PricingHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/position/category",
					Handler: positionCategoryHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/position/query",
					Handler: positionQueryHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/position/price",
					Handler: positionPriceHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/position/element",
					Handler: positionElementHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/position/placement",
					Handler: positionPlacementHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/asset/sync",
					Handler: syncAssetHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/asset/list",
					Handler: assetListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/asset/delete",
					Handler: assetDeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/asset/upload",
					Handler: assetUploadHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/asset/token",
					Handler: assetFileTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/asset/dimension",
					Handler: assetDimensionHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/asset/bind",
					Handler: assetBindHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/marketing"),
	)
}
