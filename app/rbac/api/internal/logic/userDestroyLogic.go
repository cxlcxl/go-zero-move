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

type UserDestroyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDestroyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDestroyLogic {
	return &UserDestroyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDestroyLogic) UserDestroy(req *types.UserDestroyReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.RbacClient.UserDestroy(l.ctx, &rbaccenter.UserDestroyReq{
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
