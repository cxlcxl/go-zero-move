package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

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

func (l *AccountSearchLogic) AccountSearch(in *marketing.AccountSearchReq) (*marketing.AccountSearchResp, error) {
	if in.AccountName == "" {
		return &marketing.AccountSearchResp{
			Accounts: nil,
		}, nil
	}
	accounts, err := l.svcCtx.AccountModel.RemoteAccounts(l.ctx, in.AccountName)
	if err != nil {
		return nil, err
	}
	var res []*marketing.AccountInfo
	err = copier.Copy(&res, accounts)
	if err != nil {
		return nil, err
	}
	return &marketing.AccountSearchResp{
		Accounts: res,
	}, nil
}
