package svc

import (
	"business/app/application/api/internal/config"
	"business/app/application/api/internal/middleware"
	"business/app/application/rpc/appcenter"
	"business/app/marketing/rpc/marketingcenter"
	"business/common/validators"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config               config.Config
	AuthMiddleware       rest.Middleware
	PermissionMiddleware rest.Middleware
	Validate             *validator.Validate
	AppRpcClient         appcenter.AppCenter
	AccountRpcClient     marketingcenter.MarketingCenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:               c,
		AuthMiddleware:       middleware.NewAuthMiddleware().Handle,
		PermissionMiddleware: middleware.NewPermissionMiddleware().Handle,
		Validate:             validators.New(),
		AppRpcClient:         appcenter.NewAppCenter(zrpc.MustNewClient(c.AppRpcConf)),
		AccountRpcClient:     marketingcenter.NewMarketingCenter(zrpc.MustNewClient(c.AccountRpcConf)),
	}
}
