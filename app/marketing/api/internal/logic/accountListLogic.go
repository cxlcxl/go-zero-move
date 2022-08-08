package logic

import (
	"business/app/marketing/rpc/marketing"
	"business/common/utils"
	"business/common/vars"
	"context"

	"github.com/jinzhu/copier"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountListLogic {
	return &AccountListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountListLogic) AccountList(req *types.AccountListReq) (resp *types.AccountListResp, err error) {
	if err = l.svcCtx.Validator.StructCtx(l.ctx, req); err != nil {
		return nil, err
	}
	list, err := l.svcCtx.MarketingRpcClient.AccountList(l.ctx, &marketing.AccountListReq{
		AccountName: req.AccountName,
		AccountType: req.AccountType,
		State:       req.State,
		Page:        req.Page,
		PageSize:    req.PageSize,
	})
	if err != nil {
		return nil, utils.RpcError(err, "未查询到数据")
	}
	var accounts []types.AccountInfo
	err = copier.Copy(&accounts, list.Accounts)
	if err != nil {
		return nil, err
	}
	return &types.AccountListResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Total:        list.Total,
		Data:         accounts,
		AccountTypes: vars.AccountType,
		State:        vars.CommonState,
	}, nil
}
