package svc

import (
	"business/app/marketing/api/internal/config"
	"business/app/marketing/api/internal/middleware"
	"business/app/marketing/rpc/marketingcenter"
	"business/common/validators"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config               config.Config
	AuthMiddleware       rest.Middleware
	PermissionMiddleware rest.Middleware
	Validator            *validator.Validate
	MarketingRpcClient   marketingcenter.MarketingCenter
	RedisCache           *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:               c,
		AuthMiddleware:       middleware.NewAuthMiddleware().Handle,
		PermissionMiddleware: middleware.NewPermissionMiddleware().Handle,
		Validator:            validators.New(),
		MarketingRpcClient:   marketingcenter.NewMarketingCenter(zrpc.MustNewClient(c.MarketingRpcClient)),
		RedisCache:           redis.New(c.Redis.Host),
	}
}
