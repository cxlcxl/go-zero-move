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

func (l *SetTokenLogic) SetToken(in *marketing.TokenInfo) (*marketing.BaseResp, error) {
	token, err := l.svcCtx.TokenModel.FindOneByClientId(l.ctx, in.ClientId)
	if err != nil {
		return nil, err
	}
	expire := time.Now().Unix() + in.ExpiredAt - 20
	if token != nil {
		err = l.svcCtx.TokenModel.Update(l.ctx, &model.Tokens{
			Id:           token.Id,
			ClientId:     in.ClientId,
			AccessToken:  in.AccessToken,
			RefreshToken: in.RefreshToken,
			ExpiredAt:    time.Unix(expire, 0),
			UpdatedAt:    time.Now(),
			TokenType:    in.TokenType,
		})
	} else {
		_, err = l.svcCtx.TokenModel.Insert(l.ctx, &model.Tokens{
			ClientId:     in.ClientId,
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
	_ = l.svcCtx.AccountModel.SetIsAuth(l.ctx, in.ClientId)
	return &marketing.BaseResp{
		Code: vars.ResponseCodeOk,
	}, nil
}
