package logic

import (
	"business/app/marketing/rpc/model"
	"business/common/vars"
	"context"
	"time"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

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

func (l *AccountCreateLogic) AccountCreate(in *marketing.AccountCreateReq) (*marketing.BaseResp, error) {
	now := time.Now()
	_, err := l.svcCtx.AccountModel.Insert(l.ctx, &model.Accounts{
		AdvertiserId: in.AdvertiserId,
		DeveloperId:  in.DeveloperId,
		AccountType:  in.AccountType,
		State:        1,
		AccountName:  in.AccountName,
		ClientId:     in.ClientId,
		Secret:       in.Secret,
		CreatedAt:    now,
		UpdatedAt:    now,
	})
	if err != nil {
		return nil, err
	}
	return &marketing.BaseResp{
		Code: vars.ResponseCodeOk,
	}, nil
}
