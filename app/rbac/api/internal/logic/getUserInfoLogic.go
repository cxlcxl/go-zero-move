package logic

import (
	"business/app/rbac/rpc/rbaccenter"
	"business/common/utils"
	"business/common/vars"
	"context"
	"errors"
	"github.com/jinzhu/copier"

	"business/app/rbac/api/internal/svc"
	"business/app/rbac/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	if req.Id <= 0 {
		return nil, errors.New("请求错误")
	}
	user, err := l.svcCtx.RbacClient.GetUserInfoById(l.ctx, &rbaccenter.UserInfoIdReq{Id: req.Id})
	if err != nil {
		return nil, utils.RpcError(err, "用户不存在")
	}
	var userInfo types.UserInfo
	err = copier.Copy(&userInfo, user)
	if err != nil {
		return nil, vars.DataCopierError
	}
	return &types.UserInfoResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: userInfo,
	}, nil
}
