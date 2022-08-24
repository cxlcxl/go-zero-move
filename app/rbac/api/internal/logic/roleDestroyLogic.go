package logic

import (
	"business/app/rbac/rpc/rbaccenter"
	"business/common/utils"
	"business/common/vars"
	"context"

	"business/app/rbac/api/internal/svc"
	"business/app/rbac/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDestroyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleDestroyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDestroyLogic {
	return &RoleDestroyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleDestroyLogic) RoleDestroy(req *types.RoleDestroyReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.RbacClient.RoleDestroy(l.ctx, &rbaccenter.RoleDestroyReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, utils.RpcError(err, "")
	}
	return &types.BaseResp{
		Code: vars.ResponseCodeOk,
		Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
	}, nil
}
