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

type GetAccountInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAccountInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountInfoLogic {
	return &GetAccountInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAccountInfoLogic) GetAccountInfo(req *types.AccountInfoReq) (resp *types.GetAccountInfoResp, err error) {
	info, err := l.svcCtx.AccountRpcClient.ActInfo(l.ctx, &accountcenter.ActInfoReq{Id: req.Id})
	if err != nil {
		return nil, utils.RpcError(err, "未查询到账户信息")
	}
	var actInfo types.AccountInfo
	err = copier.Copy(&actInfo, info.Info)
	if err != nil {
		return nil, vars.DataCopierError
	}
	return &types.GetAccountInfoResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: actInfo,
	}, nil
}
