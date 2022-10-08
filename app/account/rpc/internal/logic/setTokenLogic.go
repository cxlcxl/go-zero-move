package logic

import (
	"business/app/account/rpc/model"
	"context"
	"errors"
	"time"

	"business/app/account/rpc/account"
	"business/app/account/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetTokenLogic {
	return &SetTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetTokenLogic) SetToken(in *account.TokenInfo) (*account.BaseResp, error) {
	token, err := l.svcCtx.TokenModel.FindOneByAccountId(l.ctx, in.AccountId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, err
	}
	expire := time.Now().Unix() + in.ExpiredAt - 20
	if token != nil {
		err = l.svcCtx.TokenModel.Update(l.ctx, &model.Tokens{
			Id:           token.Id,
			AccountId:    in.AccountId,
			AdvertiserId: token.AdvertiserId,
			AccessToken:  in.AccessToken,
			RefreshToken: in.RefreshToken,
			ExpiredAt:    time.Unix(expire, 0),
			UpdatedAt:    time.Now(),
			TokenType:    in.TokenType,
		})
	} else {
		act, err := l.svcCtx.AccountModel.FindOne(l.ctx, in.AccountId)
		if err != nil {
			return nil, err
		}
		_, err = l.svcCtx.TokenModel.Insert(l.ctx, &model.Tokens{
			AccountId:    in.AccountId,
			AdvertiserId: act.AdvertiserId,
			AccessToken:  in.AccessToken,
			RefreshToken: in.RefreshToken,
			ExpiredAt:    time.Unix(expire, 0),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			TokenType:    in.TokenType,
		})
	}
	if err != nil {
		return nil, err
	}
	_ = l.svcCtx.AccountModel.SetIsAuth(l.ctx, in.AccountId)
	return &account.BaseResp{}, nil
}
