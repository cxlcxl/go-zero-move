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

type RoleUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateLogic {
	return &RoleUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleUpdateLogic) RoleUpdate(in *rbac.RoleUpdateReq) (*rbac.BaseResp, error) {
	if err := l.svcCtx.RoleModel.Update(l.ctx, &model.Roles{RoleName: in.RoleName, Id: in.Id, State: in.State, UpdatedAt: time.Now()}); err != nil {
		return nil, err
	}
	return &rbac.BaseResp{Code: vars.ResponseCodeOk}, nil
}
