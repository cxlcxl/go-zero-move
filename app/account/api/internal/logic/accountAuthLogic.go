package logic

import (
	"business/app/account/rpc/accountcenter"
	"business/common/curl"
	"business/common/utils"
	"business/common/vars"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"business/app/account/api/internal/svc"
	"business/app/account/api/internal/types"

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
	info, err := l.svcCtx.AccountRpcClient.GetAccountInfo(l.ctx, &accountcenter.AccountInfoReq{Id: req.Id})
	if err != nil {
		return nil, utils.RpcError(err, "请求有误")
	}
	clientId, secret := info.ClientId, info.Secret
	if clientId == "" || secret == "" {
		if info.ParentId == 0 {
			return nil, errors.New("请先完整填写 ClientId 与 Secret")
		} else {
			parent, err := l.svcCtx.AccountRpcClient.GetAccountInfo(l.ctx, &accountcenter.AccountInfoReq{Id: info.ParentId})
			if err != nil {
				return nil, utils.RpcError(err, "上级信息查询有误")
			}
			if parent.ClientId == "" || parent.Secret == "" {
				return nil, errors.New("上级 ClientId 与 Secret 信息也未填写")
			}
			clientId, secret = parent.ClientId, parent.Secret
		}
	}
	url, err := l.authorizeCodeUrl(info.Id, clientId, secret)
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

func (l *AccountAuthLogic) authorizeCodeUrl(id int64, clientId, secret string) (url string, err error) {
	token, err := l.svcCtx.AccountRpcClient.GetToken(l.ctx, &accountcenter.GetTokenReq{AccountId: id})
	if err != nil && !utils.IsErrNotFound(err) {
		return "", utils.RpcError(err, "")
	}
	if token != nil && time.Now().Before(time.Unix(token.ExpiredAt-300, 0)) {
		return "", errors.New("TOKEN 尚未过期，无需重新认证")
	}
	state := utils.MD5(fmt.Sprintf("%d-%s-%s", id, clientId, time.Now().String()))
	baseUrl := l.svcCtx.Config.MarketingApis.Authorize.CodeUrl
	if !strings.HasSuffix(baseUrl, "?") {
		baseUrl += "?"
	}
	params := curl.HttpBuildQuery(map[string]string{
		"response_type": "code",
		"access_type":   "offline",
		"client_id":     clientId,
		"state":         state,
		"redirect_uri":  l.svcCtx.Config.MarketingApis.Authorize.RedirectUri,
		"scope":         l.svcCtx.Config.MarketingApis.Authorize.Scope,
	})
	authorizeValue := fmt.Sprintf("%d-%s-%s", id, clientId, secret)
	err = l.svcCtx.RedisCache.SetCtx(l.ctx, vars.SysCachePrefix+"authorize:"+state, authorizeValue)
	if err != nil {
		return "", err
	}
	_ = l.svcCtx.RedisCache.Expire(vars.SysCachePrefix+"authorize:"+state, 300) // 保存 5 分钟
	return baseUrl + params, nil
}
