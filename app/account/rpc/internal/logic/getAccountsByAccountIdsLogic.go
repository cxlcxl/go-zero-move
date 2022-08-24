package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"business/app/account/rpc/account"
	"business/app/account/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccountsByAccountIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAccountsByAccountIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountsByAccountIdsLogic {
	return &GetAccountsByAccountIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAccountsByAccountIdsLogic) GetAccountsByAccountIds(in *account.GetByAccountIdsReq) (*account.AccountSearchResp, error) {
	accounts, err := l.svcCtx.AccountModel.GetAccountsByIds(l.ctx, in.Ids)
	if err != nil {
		return nil, err
	}
	var tmp []*account.AccountInfo
	if err = copier.Copy(&tmp, accounts); err != nil {
		return nil, err
	}
	return &account.AccountSearchResp{
		Accounts: tmp,
	}, nil
}
