package logic

import (
	"business/app/account/rpc/model"
	"business/common/utils"
	"business/common/vars"
	"context"
	"github.com/jinzhu/copier"

	"business/app/account/rpc/account"
	"business/app/account/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActListLogic {
	return &ActListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActListLogic) ActList(in *account.ActListReq) (*account.ActListResp, error) {
	offset, limit := utils.Pagination(in.Page, in.PageSize)
	list, total, err := l.svcCtx.ActModel.AccountList(l.ctx, in.AccountType, in.State, in.AccountName, in.AccountId, offset, limit)
	if err != nil {
		return nil, err
	}
	accounts := make([]*account.ActInfo, 0)
	if total > 0 {
		accountIds := make([]string, 0)
		for _, a := range list {
			accounts = append(accounts, &account.ActInfo{
				Id:               a.Id,
				AccountName:      a.AccountName,
				AccountId:        a.AccountId,
				AccountType:      a.AccountType,
				Email:            a.Email,
				State:            a.State,
				CreatedAt:        a.CreatedAt.Unix(),
				UpdatedAt:        a.UpdatedAt.Unix(),
				AccountClientIds: nil,
			})
			accountIds = append(accountIds, a.AccountId)
		}
		clientIds, err := l.svcCtx.ActClientModel.GetClientIdsByAccountIds(l.ctx, accountIds)
		if err != nil {
			return nil, err
		}
		if len(clientIds) > 0 {
			tmp := make(map[string][]*model.AccountClientIds)
			for _, clientId := range clientIds {
				tmp[clientId.AccountId] = append(tmp[clientId.AccountId], clientId)
			}
			for i, act := range accounts {
				if actClientIds, ok := tmp[act.AccountId]; ok {
					err = copier.Copy(&accounts[i].AccountClientIds, actClientIds)
					if err != nil {
						return nil, vars.DataCopierError
					}
				}
			}
		}
	}
	return &account.ActListResp{
		Total:    total,
		Accounts: accounts,
	}, nil
}
