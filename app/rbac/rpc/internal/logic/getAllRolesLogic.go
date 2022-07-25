package logic

import (
	"business/app/rbac/rpc/internal/svc"
	"business/app/rbac/rpc/rbac"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllRolesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllRolesLogic {
	return &GetAllRolesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllRolesLogic) GetAllRoles(in *rbac.AllRoles) (*rbac.AllRolesResp, error) {
	list, err := l.svcCtx.RoleModel.RoleList(l.ctx, in.RoleName, in.State, []int64{})
	if err != nil {
		return nil, err
	}
	roles := make([]*rbac.RoleDetail, len(list))
	for i := 0; i < len(list); i++ {
		roles[i] = &rbac.RoleDetail{
			Id:       list[i].Id,
			RoleName: list[i].RoleName,
			State:    list[i].State,
		}
	}
	return &rbac.AllRolesResp{
		Roles: roles,
	}, nil
}
