package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

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

func (l *GetAccountsByAccountIdsLogic) GetAccountsByAccountIds(in *marketing.GetByAccountIdsReq) (*marketing.AccountSearchResp, error) {
	accounts, err := l.svcCtx.AccountModel.GetAccountsByIds(l.ctx, in.Ids)
	if err != nil {
		return nil, err
	}
	var tmp []*marketing.AccountInfo
	if err = copier.Copy(&tmp, accounts); err != nil {
		return nil, err
	}
	return &marketing.AccountSearchResp{
		Accounts: tmp,
	}, nil
}
