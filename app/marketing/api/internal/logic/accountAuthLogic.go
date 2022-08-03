package logic

import (
	"business/app/marketing/rpc/marketingcenter"
	"business/app/marketing/rpc/model"
	"business/common/curl"
	"business/common/utils"
	"business/common/vars"
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountAuthLogic {
	return &AccountAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountAuthLogic) AccountAuth(req *types.AccountInfoReq) (resp *types.AuthResp, err error) {
	info, err := l.svcCtx.MarketingRpcClient.GetAccountInfo(l.ctx, &marketingcenter.AccountInfoReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, utils.RpcError(err, "请求有误")
	}
	if info.ClientId == 0 || info.Secret == "" {
		return nil, errors.New("请先完整填写 ClientId 与 Secret")
	}
	url, err := l.authorizeCodeUrl(info.ClientId)
	if err != nil {
		return nil, err
	}
	return &types.AuthResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: url,
	}, nil
}

func (l *AccountAuthLogic) authorizeCodeUrl(clientId int64) (url string, err error) {
	token, err := l.svcCtx.MarketingRpcClient.GetToken(l.ctx, &marketingcenter.GetTokenReq{ClientId: clientId})
	if err != nil && err != model.ErrNotFound {
		return "", utils.RpcError(err, "")
	}
	if token != nil && time.Now().Before(time.Unix(token.ExpiredAt-300, 0)) {
		return "", errors.New("TOKEN 尚未过期，无需重新认证")
	}

	clientIdToString := strconv.Itoa(int(clientId))
	state := utils.MD5(clientIdToString + time.Now().String())
	baseUrl := l.svcCtx.Config.MarketingApis.Authorize.CodeUrl
	if !strings.HasSuffix(baseUrl, "?") {
		baseUrl += "?"
	}
	params := curl.HttpBuildQuery(map[string]string{
		"response_type": "code",
		"access_type":   "offline",
		"client_id":     clientIdToString,
		"state":         state,
		"redirect_uri":  l.svcCtx.Config.MarketingApis.Authorize.RedirectUri,
		"scope":         l.svcCtx.Config.MarketingApis.Authorize.Scope,
	})
	err = l.svcCtx.RedisCache.SetCtx(l.ctx, vars.SysCachePrefix+"authorize:"+state, clientIdToString)
	if err != nil {
		return "", err
	}
	_ = l.svcCtx.RedisCache.Expire(vars.SysCachePrefix+"authorize:"+state, 300) // 保存 5 分钟
	return baseUrl + params, nil
}
