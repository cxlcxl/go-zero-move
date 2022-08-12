package logic

import (
	"business/common/utils"
	"context"

	"business/app/account/rpc/account"
	"business/app/account/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAccountListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountListLogic {
	return &AccountListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AccountListLogic) AccountList(in *account.AccountListReq) (*account.AccountListResp, error) {
	offset, limit := utils.Pagination(in.Page, in.PageSize)
	list, total, err := l.svcCtx.AccountModel.AccountList(l.ctx, in.AccountType, in.State, in.AccountName, offset, limit)
	if err != nil {
		return nil, err
	}
	accounts := make([]*account.AccountInfo, 0)
	if total > 0 {
		for _, a := range list {
			accounts = append(accounts, &account.AccountInfo{
				Id:           a.Id,
				AccountName:  a.AccountName,
				AccountType:  a.AccountType,
				AdvertiserId: a.AdvertiserId,
				DeveloperId:  a.DeveloperId,
				ClientId:     a.ClientId,
				Secret:       a.Secret,
				State:        a.State,
				IsAuth:       a.IsAuth,
				ParentId:     a.ParentId,
				CreatedAt:    a.CreatedAt.Unix(),
				UpdatedAt:    a.UpdatedAt.Unix(),
			})
		}
	}
	return &account.AccountListResp{
		Total:    total,
		Accounts: accounts,
	}, nil
}
