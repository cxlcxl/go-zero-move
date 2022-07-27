package logic

import (
	"business/app/account/rpc/model"
	"business/common/vars"
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
	token, err := l.svcCtx.TokenModel.FindOneByClientId(l.ctx, in.ClientId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, err
	}
	if token != nil {
		err = l.svcCtx.TokenModel.Update(l.ctx, &model.Tokens{
			Id:           token.Id,
			ClientId:     in.ClientId,
			AccessToken:  in.AccessToken,
			RefreshToken: in.RefreshToken,
			ExpiredAt:    time.Unix(in.ExpiredAt, 0),
			UpdatedAt:    time.Now(),
		})
	} else {
		_, err = l.svcCtx.TokenModel.Insert(l.ctx, &model.Tokens{
			ClientId:     in.ClientId,
			AccessToken:  in.AccessToken,
			RefreshToken: in.RefreshToken,
			ExpiredAt:    time.Unix(in.ExpiredAt, 0),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		})
	}
	if err != nil {
		return nil, err
	}
	return &account.BaseResp{
		Code: vars.ResponseCodeOk,
	}, nil
}
