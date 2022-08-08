package logic

import (
	"business/app/marketing/rpc/marketingcenter"
	"business/common/curl"
	"business/common/utils"
	"business/common/vars"
	"context"
	"errors"
	"fmt"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

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
	if req.ClientId == "" {
		return nil, errors.New("请求有误")
	}
	token, err := l.svcCtx.MarketingRpcClient.GetToken(l.ctx, &marketingcenter.GetTokenReq{ClientId: req.ClientId})
	if err != nil {
		return nil, utils.RpcError(err, "此 ClientId 尚未进行认证，请先认证")
	}
	info, err := l.svcCtx.MarketingRpcClient.GetAccountByClientId(l.ctx, &marketingcenter.GetTokenReq{ClientId: req.ClientId})
	if err != nil {
		return nil, utils.RpcError(err, "")
	}
	data := map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": token.RefreshToken,
		"client_id":     req.ClientId,
		"client_secret": info.Secret,
	}
	post := curl.New(l.svcCtx.Config.MarketingApis.Authorize.Refresh).Post().QueryData(data)
	var at AT
	err = post.Request(&at, curl.FormHeader())
	if err != nil {
		return nil, err
	}
	if at.Error != 0 {
		return nil, errors.New("华为接口调用失败：" + at.ErrorDescription)
	}
	fmt.Println(at)
	_, err = l.svcCtx.MarketingRpcClient.SetToken(l.ctx, &marketingcenter.TokenInfo{
		ClientId:     req.ClientId,
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
