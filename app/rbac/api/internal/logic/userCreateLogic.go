package logic

import (
	"business/app/rbac/rpc/rbaccenter"
	"business/common/utils"
	"business/common/validators"
	"business/common/vars"
	"context"

	"business/app/rbac/api/internal/svc"
	"business/app/rbac/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCreateLogic) UserCreate(req *types.UserCreateReq) (resp *types.BaseResp, err error) {
	if err = validators.New(
		validators.Email(req.Email),
		validators.Password(req.Pass),
		validators.Empty(req.Username),
	); err != nil {
		return nil, err
	}
	_, err = l.svcCtx.RbacClient.UserCreate(l.ctx, &rbaccenter.UserCreateReq{
		RoleId:   req.RoleId,
		Username: req.Username,
		Email:    req.Email,
		Mobile:   req.Mobile,
		Pass:     req.Pass,
	})
	if err != nil {
		return nil, utils.RpcError(err, "")
	}
	return &types.BaseResp{
		Code: vars.ResponseCodeOk,
		Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
	}, nil
}
