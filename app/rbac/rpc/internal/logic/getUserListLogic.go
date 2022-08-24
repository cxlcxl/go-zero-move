package logic

import (
	"business/app/rbac/rpc/internal/svc"
	"business/app/rbac/rpc/rbac"
	"business/common/utils"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserListLogic) GetUserList(in *rbac.UserListReq) (*rbac.UserListResp, error) {
	offset, limit := utils.Pagination(in.Page, in.PageSize)
	list, total, err := l.svcCtx.UserModel.UserList(l.ctx, in.Username, in.Email, in.State, offset, limit)
	if err != nil {
		return nil, err
	}
	users := make([]*rbac.UserInfoResp, len(list))
	roleIds := make([]int64, 0)
	for i := range list {
		roleIds = append(roleIds, list[i].RoleId)
	}
	roleList, err := l.svcCtx.RoleModel.RoleList(l.ctx, "", -1, utils.ArrayUnique(roleIds))
	if err != nil {
		return nil, err
	}
	tmp := make(map[int64]*rbac.RoleDetail)
	for i := range roleList {
		tmp[roleList[i].Id] = &rbac.RoleDetail{
			Id:       roleList[i].Id,
			RoleName: roleList[i].RoleName,
			State:    roleList[i].State,
		}
	}
	for i := range list {
		role := tmp[list[i].RoleId]
		users[i] = &rbac.UserInfoResp{
			Id:        list[i].Id,
			Username:  list[i].Username,
			Email:     list[i].Email,
			Mobile:    list[i].Mobile,
			State:     list[i].State,
			CreatedAt: list[i].CreatedAt.Unix(),
			Role:      role,
		}
	}
	return &rbac.UserListResp{
		Total: total,
		Users: users,
	}, nil
}
