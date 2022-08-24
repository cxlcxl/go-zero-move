package logic

import (
	"business/app/rbac/rpc/internal/svc"
	"business/app/rbac/rpc/rbac"
	"business/common/vars"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDestroyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleDestroyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDestroyLogic {
	return &RoleDestroyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleDestroyLogic) RoleDestroy(in *rbac.RoleDestroyReq) (*rbac.BaseResp, error) {
	exists, err := l.svcCtx.UserModel.FindExistsUsersByRoleId(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("角色已绑定账号，不能删除")
	}
	err = l.svcCtx.RoleModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &rbac.BaseResp{Code: vars.ResponseCodeOk}, nil
}
