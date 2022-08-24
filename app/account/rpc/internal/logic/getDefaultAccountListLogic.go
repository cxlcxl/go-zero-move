package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"business/app/account/rpc/account"
	"business/app/account/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDefaultAccountListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDefaultAccountListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDefaultAccountListLogic {
	return &GetDefaultAccountListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDefaultAccountListLogic) GetDefaultAccountList(in *account.DefaultListReq) (*account.AccountSearchResp, error) {
	accounts, err := l.svcCtx.AccountModel.RemoteAccounts(l.ctx, "", 0)
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
