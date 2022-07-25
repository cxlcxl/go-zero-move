package logic

import (
	"context"

	"business/app/rbac/rpc/internal/svc"
	"business/app/rbac/rpc/rbac"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByEmailLogic {
	return &GetUserInfoByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoByEmailLogic) GetUserInfoByEmail(in *rbac.UserInfoEmailReq) (*rbac.UserInfoResp, error) {
	// todo: add your logic here and delete this line

	return &rbac.UserInfoResp{}, nil
}
