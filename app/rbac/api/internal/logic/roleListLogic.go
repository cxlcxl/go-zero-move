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

type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleListLogic) RoleList(req *types.RoleListReq) (resp *types.RoleListResp, err error) {
	roles, err := l.svcCtx.RbacClient.GetAllRoles(l.ctx, &rbaccenter.AllRoles{
		RoleName: req.RoleName,
		State:    req.State,
	})
	if err != nil {
		return nil, utils.RpcError(err, "没有相关数据")
	}
	var roleList []types.RoleInfo
	err = copier.Copy(&roleList, roles.Roles)
	if err != nil {
		return nil, vars.DataCopierError
	}
	return &types.RoleListResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: roleList,
	}, nil
}
