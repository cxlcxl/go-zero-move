package logic

import (
	"business/app/rbac/api/internal/svc"
	"business/app/rbac/api/internal/types"
	"business/app/rbac/rpc/rbaccenter"
	"business/common/utils"
	"business/common/vars"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type RoleUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateLogic {
	return &RoleUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleUpdateLogic) RoleUpdate(req *types.RoleUpdateReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.RbacClient.RoleUpdate(l.ctx, &rbaccenter.RoleUpdateReq{
		Id:       req.Id,
		RoleName: req.RoleName,
		State:    req.State,
	})
	if err != nil {
		return nil, utils.RpcError(err, "")
	}

	return &types.BaseResp{
		Code: vars.ResponseCodeOk,
		Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
	}, nil
}
