package logic

import (
	"context"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccountSecretByClientIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAccountSecretByClientIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountSecretByClientIdLogic {
	return &GetAccountSecretByClientIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAccountSecretByClientIdLogic) GetAccountSecretByClientId(in *marketing.GetTokenReq) (*marketing.AccountInfo, error) {
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
