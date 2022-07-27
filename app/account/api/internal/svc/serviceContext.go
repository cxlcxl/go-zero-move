package svc

import (
	"business/app/account/api/internal/config"
	"business/app/account/api/internal/middleware"
	"business/app/account/rpc/accountcenter"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config               config.Config
	AuthMiddleware       rest.Middleware
	PermissionMiddleware rest.Middleware
	AccountRpcClient     accountcenter.AccountCenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:               c,
		AuthMiddleware:       middleware.NewAuthMiddleware().Handle,
		PermissionMiddleware: middleware.NewPermissionMiddleware().Handle,
		AccountRpcClient:     accountcenter.NewAccountCenter(zrpc.MustNewClient(c.AccountRpcConf)),
	}
}
