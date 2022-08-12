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

type AccountDefaultListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountDefaultListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountDefaultListLogic {
	return &AccountDefaultListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountDefaultListLogic) AccountDefaultList() (resp *types.AccountSearchResp, err error) {
	var res []types.AccountInfo
	accounts, err := l.svcCtx.AccountRpcClient.GetDefaultAccountList(l.ctx, &accountcenter.DefaultListReq{})
	if err != nil {
		return nil, utils.RpcError(err, "未查到账户数据")
	}
	if err = copier.Copy(&res, accounts.Accounts); err != nil {
		return nil, vars.DataCopierError
	}

	return &types.AccountSearchResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: res,
	}, nil
}
