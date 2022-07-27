package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"business/app/account/rpc/account"
	"business/app/account/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActInfoLogic {
	return &ActInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActInfoLogic) ActInfo(in *account.ActInfoReq) (*account.ActInfoResp, error) {
	accountInfo, err := l.svcCtx.ActModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	ids, err := l.svcCtx.ActClientModel.GetClientIdsByAccountIds(l.ctx, []string{accountInfo.AccountId})
	if err != nil {
		return nil, err
	}
	var clientIds []*account.ActClientId
	err = copier.Copy(&clientIds, ids)
	if err != nil {
		return nil, err
	}
	return &account.ActInfoResp{
		Info: &account.ActInfo{
			Id:               accountInfo.Id,
			AccountName:      accountInfo.AccountName,
			AccountId:        accountInfo.AccountId,
			AccountType:      accountInfo.AccountType,
			Email:            accountInfo.Email,
			State:            accountInfo.State,
			CreatedAt:        accountInfo.CreatedAt.Unix(),
			UpdatedAt:        accountInfo.UpdatedAt.Unix(),
			AccountClientIds: clientIds,
		},
	}, nil
}
