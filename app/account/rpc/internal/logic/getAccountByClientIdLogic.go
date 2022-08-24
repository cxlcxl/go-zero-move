package logic

import (
	"context"

	"business/app/account/rpc/account"
	"business/app/account/rpc/internal/svc"

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

func (l *GetAccountByClientIdLogic) GetAccountByClientId(in *account.GetTokenReq) (*account.AccountInfo, error) {
	info, err := l.svcCtx.AccountModel.FindOne(l.ctx, in.AccountId)
	if err != nil {
		return nil, err
	}
	return &account.AccountInfo{
		Id:           info.Id,
		AccountName:  info.AccountName,
		AccountType:  info.AccountType,
		AdvertiserId: info.AdvertiserId,
		DeveloperId:  info.DeveloperId,
		ClientId:     info.ClientId,
		Secret:       info.Secret,
		State:        info.State,
		ParentId:     info.ParentId,
		CreatedAt:    info.CreatedAt.Unix(),
		UpdatedAt:    info.UpdatedAt.Unix(),
	}, nil
}
