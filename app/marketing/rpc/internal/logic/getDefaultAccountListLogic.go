package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

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

func (l *GetDefaultAccountListLogic) GetDefaultAccountList(in *marketing.DefaultListReq) (*marketing.AccountSearchResp, error) {
	accounts, err := l.svcCtx.AccountModel.RemoteAccounts(l.ctx, "")
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
