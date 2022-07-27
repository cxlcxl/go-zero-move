package logic

import (
	"business/app/account/rpc/account"
	"business/app/account/rpc/internal/svc"
	"context"

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

func (l *GetTokenLogic) GetToken(in *account.GetTokenReq) (*account.TokenInfo, error) {
	token, err := l.svcCtx.TokenModel.FindOneByClientId(l.ctx, in.ClientId)
	if err != nil {
		return nil, err
	}
	return &account.TokenInfo{
		ClientId:     token.ClientId,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiredAt:    token.ExpiredAt.Unix(),
	}, nil
}
