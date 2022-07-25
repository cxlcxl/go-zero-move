package logic

import (
	"business/app/rbac/rpc/internal/svc"
	"business/app/rbac/rpc/model"
	"business/app/rbac/rpc/rbac"
	"business/common/vars"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type RoleCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleCreateLogic {
	return &RoleCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleCreateLogic) RoleCreate(in *rbac.RoleCreateReq) (*rbac.BaseResp, error) {
	_, err := l.svcCtx.RoleModel.Insert(l.ctx, &model.Roles{
		RoleName:  in.RoleName,
		State:     1,
		Sys:       in.Sys,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &rbac.BaseResp{Code: vars.ResponseCodeOk}, nil
}
