package logic

import (
	"business/app/account/api/internal/svc"
	"business/app/account/api/internal/types"
	"business/app/account/rpc/accountcenter"
	"business/common/curl"
	"business/common/utils"
	"business/common/vars"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.RefreshTokenReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.AccountRpcClient.RefreshToken(l.ctx, &accountcenter.GetTokenReq{AccountId: req.Id})
	if err != nil {
		return nil, utils.RpcError(err, "请检查用户是否填写完整 ClientId 信息")
	}

	return &types.BaseResp{
		Code: vars.ResponseCodeOk,
		Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
	}, nil
}

func (l *RefreshTokenLogic) refreshBak(req *types.RefreshTokenReq) (resp *types.BaseResp, err error) {
	token, err := l.svcCtx.AccountRpcClient.GetToken(l.ctx, &accountcenter.GetTokenReq{AccountId: req.Id})
	if err != nil {
		return nil, utils.RpcError(err, "此 ClientId 尚未进行认证，请先认证")
	}
	info, err := l.svcCtx.AccountRpcClient.GetAccountInfo(l.ctx, &accountcenter.AccountInfoReq{Id: req.Id})
	if err != nil {
		return nil, utils.RpcError(err, "")
	}
	clientId, secret := info.ClientId, info.Secret
	if clientId == "" || secret == "" {
		if info.ParentId == 0 {
			return nil, errors.New("请先完整填写 ClientId 与 Secret")
		} else {
			parent, err := l.svcCtx.AccountRpcClient.GetAccountInfo(l.ctx, &accountcenter.AccountInfoReq{Id: info.ParentId})
			if err != nil || parent.ClientId == "" || parent.Secret == "" {
				return nil, utils.RpcError(err, "检查上级 ClientId 信息有误，请检查是否有完整填写")
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
	var at AT
	if err = curl.New(l.svcCtx.Config.MarketingApis.Authorize.Refresh).Post().QueryData(data).Request(&at, curl.FormHeader()); err != nil {
		return nil, err
	}
	if at.Error != 0 {
		return nil, errors.New("华为接口调用失败：" + at.ErrorDescription)
	}
	_, err = l.svcCtx.AccountRpcClient.SetToken(l.ctx, &accountcenter.TokenInfo{
		AccountId:    req.Id,
		AccessToken:  at.AccessToken,
		RefreshToken: at.RefreshToken,
		ExpiredAt:    at.ExpiresIn,
		TokenType:    at.TokenType,
	})
	if err != nil {
		return nil, utils.RpcError(err, "")
	}
	return &types.BaseResp{
		Code: vars.ResponseCodeOk,
		Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
	}, nil
}
