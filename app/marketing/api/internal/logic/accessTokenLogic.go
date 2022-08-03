package logic

import (
	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"
	"business/app/marketing/rpc/marketingcenter"
	"business/common/curl"
	"business/common/utils"
	"business/common/vars"
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccessTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccessTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccessTokenLogic {
	return &AccessTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type AT struct {
	Error            int64  `json:"error"`
	ErrorDescription string `json:"error_description"`
	AccessToken      string `json:"access_token"`
	ExpiresIn        int64  `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
}

func (l *AccessTokenLogic) AccessToken(req *types.AccessTokenReq) (resp *types.AccessTokenResp, err error) {
	if err = l.svcCtx.Validator.StructCtx(l.ctx, req); err != nil {
		return nil, err
	}
	clientId, err := l.svcCtx.RedisCache.Get(vars.SysCachePrefix + "authorize:" + req.State)
	if err != nil {
		return nil, err
	}
	_, _ = l.svcCtx.RedisCache.Del(vars.SysCachePrefix + "authorize:" + req.State)
	clientIdToInt, err := strconv.ParseInt(clientId, 0, 64)
	if err != nil {
		return nil, errors.New("参数已过期，请重新发起认证")
	}
	info, err := l.svcCtx.MarketingRpcClient.GetAccountSecretByClientId(l.ctx, &marketingcenter.GetTokenReq{
		ClientId: clientIdToInt,
	})
	if err != nil {
		return nil, utils.RpcError(err, "请求错误")
	}
	data := map[string]string{
		"grant_type":    "authorization_code",
		"code":          req.AuthorizationCode,
		"client_id":     clientId,
		"client_secret": info.Secret,
		"redirect_uri":  l.svcCtx.Config.MarketingApis.Authorize.RedirectUri,
	}
	post := curl.New(l.svcCtx.Config.MarketingApis.Authorize.TokenUrl).Post().QueryData(data)
	var at AT
	if err = post.Request(&at, curl.FormHeader()); err != nil {
		return nil, err
	}
	if at.Error != 0 {
		return nil, errors.New("华为接口调用失败：" + at.ErrorDescription)
	}
	fmt.Println(at)
	_, err = l.svcCtx.MarketingRpcClient.SetToken(l.ctx, &marketingcenter.TokenInfo{
		ClientId:     clientIdToInt,
		AccessToken:  at.AccessToken,
		RefreshToken: at.RefreshToken,
		ExpiredAt:    at.ExpiresIn,
		TokenType:    at.TokenType,
	})
	if err != nil {
		return nil, utils.RpcError(err, "")
	}
	return &types.AccessTokenResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		AccessToken: at.AccessToken,
	}, nil
}
