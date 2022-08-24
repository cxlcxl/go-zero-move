package logic

import (
	"business/app/account/rpc/accountcenter"
	"business/common/utils"
	"business/common/vars"
	"context"

	"github.com/jinzhu/copier"

	"business/app/account/api/internal/svc"
	"business/app/account/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountInfoLogic {
	return &AccountInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountInfoLogic) AccountInfo(req *types.AccountInfoReq) (resp *types.AccountInfoResp, err error) {
	info, err := l.svcCtx.AccountRpcClient.GetAccountInfo(l.ctx, &accountcenter.AccountInfoReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, utils.RpcError(err, "请求有误")
	}

	var act types.AccountInfo
	err = copier.Copy(&act, info)
	if err != nil {
		return nil, err
	}
	return &types.AccountInfoResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: act,
	}, nil
}
