package logic

import (
	"business/app/account/rpc/model"
	"business/common/curl"
	"business/common/vars"
	"context"
	"errors"
	"time"

	"business/app/account/rpc/account"
	"business/app/account/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshTokenLogic) RefreshToken(in *account.GetTokenReq) (*account.TokenInfo, error) {
	token, err := l.svcCtx.TokenModel.FindOneByAccountId(l.ctx, in.AccountId)
	if err != nil {
		return nil, err
	}
	info, err := l.svcCtx.AccountModel.FindOne(l.ctx, in.AccountId)
	if err != nil {
		return nil, err
	}
	clientId, secret := info.ClientId, info.Secret
	if clientId == "" || secret == "" {
		if info.ParentId == 0 {
			return nil, errors.New("请先完整填写 ClientId 与 Secret")
		} else {
			parent, err := l.svcCtx.AccountModel.FindOne(l.ctx, info.ParentId)
			if err != nil {
				return nil, err
			}
			if parent.ClientId == "" || parent.Secret == "" {
				return nil, errors.New("检查上级 ClientId 信息有误，请检查是否有完整填写")
			}
			clientId, secret = parent.ClientId, parent.Secret
		}
	}
	data := map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": token.RefreshToken,
		"client_id":     clientId,
		"client_secret": secret,
	}
	var at vars.AdsAccessTokenResponse
	if err = curl.New(l.svcCtx.Config.MarketingApis.Refresh).Post().QueryData(data).Request(&at, curl.FormHeader()); err != nil {
		return nil, err
	}
	if at.Error != 0 {
		return nil, errors.New("华为接口调用失败：" + at.ErrorDescription)
	}
	expire := time.Now().Unix() + at.ExpiresIn - 20
	err = l.svcCtx.TokenModel.Update(l.ctx, &model.Tokens{
		Id:           token.Id,
		AccountId:    in.AccountId,
		AccessToken:  at.AccessToken,
		RefreshToken: at.RefreshToken,
		ExpiredAt:    time.Unix(expire, 0),
		UpdatedAt:    time.Now(),
		TokenType:    at.TokenType,
	})
	if err != nil {
		return nil, err
	}

	return &account.TokenInfo{
		AccountId:    in.AccountId,
		AccessToken:  at.AccessToken,
		RefreshToken: at.RefreshToken,
		ExpiredAt:    expire,
		TokenType:    at.TokenType,
	}, nil
}
