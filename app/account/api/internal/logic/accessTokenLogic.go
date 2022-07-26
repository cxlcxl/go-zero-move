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
	"strconv"
	"strings"
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
	authorizeValue, err := l.svcCtx.RedisCache.Get(vars.SysCachePrefix + "authorize:" + req.State)
	if err != nil {
		return nil, err
	}
	_, _ = l.svcCtx.RedisCache.Del(vars.SysCachePrefix + "authorize:" + req.State)
	split := strings.Split(authorizeValue, "-")
	id, err := strconv.ParseInt(split[0], 0, 64)
	if err != nil {
		return nil, err
	}

	data := map[string]string{
		"grant_type":    "authorization_code",
		"code":          req.AuthorizationCode,
		"client_id":     split[1],
		"client_secret": split[2],
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

	_, err = l.svcCtx.AccountRpcClient.SetToken(l.ctx, &accountcenter.TokenInfo{
		AccountId:    id,
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
