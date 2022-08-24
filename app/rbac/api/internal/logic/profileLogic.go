package logic

import (
	"business/app/rbac/rpc/rbaccenter"
	"business/common/utils"
	"business/common/vars"
	"context"
	"encoding/json"
	"errors"
	"github.com/jinzhu/copier"
	"strconv"

	"business/app/rbac/api/internal/svc"
	"business/app/rbac/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProfileLogic {
	return &ProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProfileLogic) Profile() (resp *types.ProfileResp, err error) {
	id, ok := l.ctx.Value("id").(json.Number)
	if !ok {
		return nil, errors.New("个人信息获取失败")
	}
	userId, err := strconv.ParseInt(string(id), 0, 64)
	if err != nil {
		return nil, errors.New("个人资料获取失败")
	}
	byId, err := l.svcCtx.RbacClient.GetUserInfoById(l.ctx, &rbaccenter.UserInfoIdReq{Id: userId})
	if err != nil {
		return nil, utils.RpcError(err, "")
	}
	var user types.UserInfo
	err = copier.Copy(&user, byId)
	if err != nil {
		return nil, vars.DataCopierError
	}

	return &types.ProfileResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: types.ProfileInfo{
			UserInfo:    user,
			Permissions: []string{},
		},
	}, nil
}
