package svc

import (
	"business/app/ads/api/internal/config"
	"business/app/ads/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config               config.Config
	AuthMiddleware       rest.Middleware
	PermissionMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:               c,
		AuthMiddleware:       middleware.NewAuthMiddleware().Handle,
		PermissionMiddleware: middleware.NewPermissionMiddleware().Handle,
	}
}
