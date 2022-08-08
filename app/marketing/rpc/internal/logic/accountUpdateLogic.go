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

func (l *AccountUpdateLogic) AccountUpdate(in *marketing.AccountUpdateReq) (*marketing.BaseResp, error) {
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
	return &marketing.BaseResp{
		Code: vars.ResponseCodeOk,
	}, nil
}
