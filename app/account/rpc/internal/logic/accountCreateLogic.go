package logic

import (
	"business/app/account/rpc/model"
	"context"
	"time"

	"business/app/account/rpc/account"
	"business/app/account/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAccountCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountCreateLogic {
	return &AccountCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AccountCreateLogic) AccountCreate(in *account.AccountCreateReq) (*account.BaseResp, error) {
	now := time.Now()
	_, err := l.svcCtx.AccountModel.Insert(l.ctx, &model.Accounts{
		AdvertiserId: in.AdvertiserId,
		DeveloperId:  in.DeveloperId,
		AccountType:  in.AccountType,
		State:        1,
		AccountName:  in.AccountName,
		ClientId:     in.ClientId,
		Secret:       in.Secret,
		ParentId:     in.ParentId,
		CreatedAt:    now,
		UpdatedAt:    now,
	})
	if err != nil {
		return nil, err
	}
	return &account.BaseResp{}, nil
}
