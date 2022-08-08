package logic

import (
	"business/app/marketing/rpc/marketingcenter"
	"business/common/utils"
	"business/common/vars"
	"context"
	"github.com/jinzhu/copier"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountSearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountSearchLogic {
	return &AccountSearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountSearchLogic) AccountSearch(req *types.AccountSearchReq) (resp *types.AccountSearchResp, err error) {
	var res []types.AccountInfo
	if req.AccountName != "" {
		accounts, err := l.svcCtx.MarketingRpcClient.AccountSearch(l.ctx, &marketingcenter.AccountSearchReq{AccountName: req.AccountName})
		if err != nil {
			return nil, utils.RpcError(err, "未查到账户数据")
		}
		if err = copier.Copy(&res, accounts.Accounts); err != nil {
			return nil, vars.DataCopierError
		}
	}
	return &types.AccountSearchResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: res,
	}, nil
}
