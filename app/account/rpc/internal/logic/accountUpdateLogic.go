package logic

import (
	"business/app/account/rpc/model"
	"context"
	"time"

	"business/app/account/rpc/account"
	"business/app/account/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAccountUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountUpdateLogic {
	return &AccountUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AccountUpdateLogic) AccountUpdate(in *account.AccountUpdateReq) (*account.BaseResp, error) {
	err := l.svcCtx.AccountModel.Update(l.ctx, &model.Accounts{
		Id:           in.Id,
		AdvertiserId: in.AdvertiserId,
		DeveloperId:  in.DeveloperId,
		AccountType:  in.AccountType,
		State:        in.State,
		AccountName:  in.AccountName,
		ClientId:     in.ClientId,
		Secret:       in.Secret,
		ParentId:     in.ParentId,
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &account.BaseResp{}, nil
}
