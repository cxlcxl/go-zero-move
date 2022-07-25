package logic

import (
	"business/app/rbac/rpc/internal/svc"
	"business/app/rbac/rpc/rbac"
	"business/common/vars"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserDestroyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDestroyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDestroyLogic {
	return &UserDestroyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserDestroyLogic) UserDestroy(in *rbac.UserDestroyReq) (*rbac.BaseResp, error) {
	if err := l.svcCtx.UserModel.Delete(l.ctx, in.Id); err != nil {
		return nil, err
	}
	return &rbac.BaseResp{Code: vars.ResponseCodeOk}, nil
}
