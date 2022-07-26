package svc

import (
	"business/app/account/api/internal/config"
	"business/app/account/api/internal/middleware"
	"business/app/account/rpc/accountcenter"
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
	AccountRpcClient     accountcenter.AccountCenter
	RedisCache           *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:               c,
		AuthMiddleware:       middleware.NewAuthMiddleware().Handle,
		PermissionMiddleware: middleware.NewPermissionMiddleware().Handle,
		Validator:            validators.New(),
		AccountRpcClient:     accountcenter.NewAccountCenter(zrpc.MustNewClient(c.AccountRpcClient)),
		RedisCache:           redis.New(c.Redis.Host),
	}
}
