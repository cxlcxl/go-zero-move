package logic

import (
	"context"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccountByClientIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAccountByClientIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountByClientIdLogic {
	return &GetAccountByClientIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAccountByClientIdLogic) GetAccountByClientId(in *marketing.GetTokenReq) (*marketing.AccountInfo, error) {
	info, err := l.svcCtx.AccountModel.FindOneByClientId(l.ctx, in.ClientId)
	if err != nil {
		return nil, err
	}
	return &marketing.AccountInfo{
		Id:           info.Id,
		AccountName:  info.AccountName,
		AccountType:  info.AccountType,
		AdvertiserId: info.AdvertiserId,
		DeveloperId:  info.DeveloperId,
		ClientId:     info.ClientId,
		Secret:       info.Secret,
		State:        info.State,
		CreatedAt:    info.CreatedAt.Unix(),
		UpdatedAt:    info.UpdatedAt.Unix(),
	}, nil
}
