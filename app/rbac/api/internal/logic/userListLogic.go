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

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.UserListReq) (resp *types.UserListResp, err error) {
	list, err := l.svcCtx.RbacClient.GetUserList(l.ctx, &rbaccenter.UserListReq{
		Username: req.Username,
		Email:    req.Email,
		State:    req.State,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, utils.RpcError(err, "未查询到相关数据")
	}
	var users []types.UserInfo
	err = copier.Copy(&users, list.Users)
	if err != nil {
		return nil, vars.DataCopierError
	}
	return &types.UserListResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Total: list.Total,
		Data:  users,
	}, nil
}
