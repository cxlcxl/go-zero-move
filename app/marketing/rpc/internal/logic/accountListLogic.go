package logic

import (
	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"
	"business/common/utils"
	"context"

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

func (l *AccountListLogic) AccountList(in *marketing.AccountListReq) (*marketing.AccountListResp, error) {
	offset, limit := utils.Pagination(in.Page, in.PageSize)
	list, total, err := l.svcCtx.AccountModel.AccountList(l.ctx, in.AccountType, in.State, in.AccountName, offset, limit)
	if err != nil {
		return nil, err
	}
	accounts := make([]*marketing.AccountInfo, 0)
	if total > 0 {
		for _, a := range list {
			accounts = append(accounts, &marketing.AccountInfo{
				Id:           a.Id,
				AccountName:  a.AccountName,
				AccountType:  a.AccountType,
				AdvertiserId: a.AdvertiserId,
				DeveloperId:  a.DeveloperId,
				ClientId:     a.ClientId,
				Secret:       a.Secret,
				State:        a.State,
				IsAuth:       a.IsAuth,
				CreatedAt:    a.CreatedAt.Unix(),
				UpdatedAt:    a.UpdatedAt.Unix(),
			})
		}
	}
	return &marketing.AccountListResp{
		Total:    total,
		Accounts: accounts,
	}, nil
}
