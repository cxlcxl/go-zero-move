package svc

import (
	"business/app/rbac/api/internal/config"
	"business/app/rbac/api/internal/middleware"
	"business/app/rbac/rpc/rbaccenter"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config               config.Config
	AuthMiddleware       rest.Middleware
	PermissionMiddleware rest.Middleware
	RbacClient           rbaccenter.RbacCenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:               c,
		AuthMiddleware:       middleware.NewAuthMiddleware().Handle,
		PermissionMiddleware: middleware.NewPermissionMiddleware().Handle,
		RbacClient:           rbaccenter.NewRbacCenter(zrpc.MustNewClient(c.RpcConf)),
	}
}
