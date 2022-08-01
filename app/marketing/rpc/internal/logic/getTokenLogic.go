package logic

import (
	"context"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenLogic {
	return &GetTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTokenLogic) GetToken(in *marketing.GetTokenReq) (*marketing.TokenInfo, error) {
	token, err := l.svcCtx.TokenModel.FindOneByClientId(l.ctx, in.ClientId)
	if err != nil {
		return nil, err
	}
	return &marketing.TokenInfo{
		ClientId:     token.ClientId,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiredAt:    token.ExpiredAt.Unix(),
		TokenType:    token.TokenType,
	}, nil
}
