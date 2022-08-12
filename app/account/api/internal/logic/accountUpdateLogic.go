package logic

import (
	"business/app/account/rpc/accountcenter"
	"business/common/utils"
	"business/common/vars"
	"context"
	"errors"

	"business/app/account/api/internal/svc"
	"business/app/account/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountUpdateLogic {
	return &AccountUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountUpdateLogic) AccountUpdate(req *types.AccountUpdateReq) (resp *types.BaseResp, err error) {
	if err = l.svcCtx.Validator.StructCtx(l.ctx, req); err != nil {
		return nil, err
	}
	if req.Id == req.ParentId {
		return nil, errors.New("上级服务商不可是本账户")
	}
	_, err = l.svcCtx.AccountRpcClient.AccountUpdate(l.ctx, &accountcenter.AccountUpdateReq{
		Id:           req.Id,
		AccountName:  req.AccountName,
		AccountType:  req.AccountType,
		AdvertiserId: req.AdvertiserId,
		DeveloperId:  req.DeveloperId,
		State:        req.State,
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
