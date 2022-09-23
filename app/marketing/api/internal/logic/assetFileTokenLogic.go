package logic

import (
	"business/app/account/rpc/accountcenter"
	"business/app/application/rpc/appcenter"
	"business/app/marketing/api/internal/statements"
	"business/common/curl"
	"business/common/utils"
	"business/common/vars"
	"context"
	"errors"
	"fmt"
	"time"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssetFileTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssetFileTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetFileTokenLogic {
	return &AssetFileTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssetFileTokenLogic) AssetFileToken(req *types.AssetFileTokenReq) (resp *types.AssetFileTokenResp, err error) {
	app, err := l.svcCtx.AppRpcClient.GetAppInfoByAppId(l.ctx, &appcenter.GetByAppIdReq{AppId: req.AppId})
	if err != nil {
		return nil, utils.RpcError(err, "应用信息有误")
	}
	token, err := l.svcCtx.AccountRpcClient.GetToken(l.ctx, &accountcenter.GetTokenReq{AccountId: app.AccountId})
	if err != nil {
		return nil, utils.RpcError(err, "Token 信息获取失败")
	}
	if time.Unix(token.ExpiredAt, 0).Before(time.Now()) {
		return nil, errors.New("App 关联账户 Token 已过期，请先到账户列表页刷新 Token ")
	}

	t := fmt.Sprintf("%s %s", token.TokenType, token.AccessToken)
	d := statements.FileTokenReq{AdvertiserId: app.AdvertiserId}
	c, err := curl.New(l.svcCtx.Config.MarketingApis.Asset.Token).Get().JsonData(d)
	if err != nil {
		return nil, err
	}
	var rs statements.FileTokenResponse
	if err = c.Request(&rs, curl.Authorization(t)); err != nil {
		return nil, err
	}
	if rs.Code != "200" {
		return nil, errors.New(rs.Message)
	}

	return &types.AssetFileTokenResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: rs.Data.FileToken,
	}, nil
}
