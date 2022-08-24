package svc

import (
	"business/app/rbac/api/internal/config"
	"business/app/rbac/api/internal/middleware"
	"business/app/rbac/rpc/rbaccenter"
	"business/common/validators"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config               config.Config
	PermissionMiddleware rest.Middleware
	RbacClient           rbaccenter.RbacCenter
	Validator            *validator.Validate
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:               c,
		PermissionMiddleware: middleware.NewPermissionMiddleware().Handle,
		RbacClient:           rbaccenter.NewRbacCenter(zrpc.MustNewClient(c.RpcConf)),
		Validator:            validators.New(),
	}
}
