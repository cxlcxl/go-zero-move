package logic

import (
	"business/app/rbac/rpc/internal/svc"
	"business/app/rbac/rpc/rbac"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleInfoLogic {
	return &GetRoleInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoleInfoLogic) GetRoleInfo(in *rbac.RoleInfo) (*rbac.RoleDetail, error) {
	role, err := l.svcCtx.RoleModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &rbac.RoleDetail{Id: role.Id, RoleName: role.RoleName}, nil
}
