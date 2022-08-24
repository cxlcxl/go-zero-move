package logic

import (
	"context"
	"errors"

	"business/app/rbac/rpc/internal/svc"
	"business/app/rbac/rpc/rbac"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByIdLogic {
	return &GetUserInfoByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoByIdLogic) GetUserInfoById(in *rbac.UserInfoIdReq) (*rbac.UserInfoResp, error) {
	if in.Id <= 0 {
		return nil, errors.New("用户信息错误")
	}
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	role, err := l.svcCtx.RoleModel.FindOne(l.ctx, user.RoleId)
	if err != nil {
		return nil, err
	}
	return &rbac.UserInfoResp{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		Mobile:    user.Mobile,
		State:     user.State,
		CreatedAt: user.CreatedAt.Unix(),
		Role: &rbac.RoleDetail{
			Id:       user.RoleId,
			RoleName: role.RoleName,
		},
	}, nil
}
