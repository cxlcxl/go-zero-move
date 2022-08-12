package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"business/app/account/rpc/account"
	"business/app/account/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountSearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAccountSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountSearchLogic {
	return &AccountSearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AccountSearchLogic) AccountSearch(in *account.AccountSearchReq) (*account.AccountSearchResp, error) {
	if in.AccountName == "" {
		return &account.AccountSearchResp{
			Accounts: nil,
		}, nil
	}
	accounts, err := l.svcCtx.AccountModel.RemoteAccounts(l.ctx, in.AccountName)
	if err != nil {
		return nil, err
	}
	var res []*account.AccountInfo
	err = copier.Copy(&res, accounts)
	if err != nil {
		return nil, err
	}
	return &account.AccountSearchResp{
		Accounts: res,
	}, nil
}
