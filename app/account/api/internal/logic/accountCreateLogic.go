package logic

import (
	"business/app/account/rpc/accountcenter"
	"business/common/utils"
	"business/common/vars"
	"context"

	"business/app/account/api/internal/svc"
	"business/app/account/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountCreateLogic {
	return &AccountCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountCreateLogic) AccountCreate(req *types.AccountPost) (resp *types.BaseResp, err error) {
	if err = l.svcCtx.Validator.StructCtx(l.ctx, req); err != nil {
		return nil, err
	}
	_, err = l.svcCtx.AccountRpcClient.AccountCreate(l.ctx, &accountcenter.AccountCreateReq{
		AccountName:  req.AccountName,
		AccountType:  req.AccountType,
		AdvertiserId: req.AdvertiserId,
		DeveloperId:  req.DeveloperId,
		ClientId:     req.ClientId,
		Secret:       req.Secret,
		ParentId:     req.ParentId,
	})
	if err != nil {
		return nil, utils.RpcError(err, "")
	}

	return &types.BaseResp{
		Code: vars.ResponseCodeOk,
		Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
	}, nil
}
