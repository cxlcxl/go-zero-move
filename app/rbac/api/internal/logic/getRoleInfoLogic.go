package logic

import (
	"business/app/rbac/rpc/rbaccenter"
	"business/common/utils"
	"business/common/vars"
	"context"
	"github.com/jinzhu/copier"

	"business/app/rbac/api/internal/svc"
	"business/app/rbac/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleInfoLogic {
	return &GetRoleInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleInfoLogic) GetRoleInfo(req *types.RoleInfoReq) (resp *types.RoleInfoResp, err error) {
	info, err := l.svcCtx.RbacClient.GetRoleInfo(l.ctx, &rbaccenter.RoleInfo{Id: req.Id})
	if err != nil {
		return nil, utils.RpcError(err, "请求错误")
	}
	var role types.RoleInfo
	err = copier.Copy(&role, info)
	if err != nil {
		return nil, vars.DataCopierError
	}
	return &types.RoleInfoResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: role,
	}, nil
}
