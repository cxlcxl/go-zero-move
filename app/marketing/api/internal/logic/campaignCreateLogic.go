package logic

import (
	"business/app/account/rpc/accountcenter"
	"business/app/application/rpc/appcenter"
	"business/app/marketing/api/internal/statements"
	"business/app/marketing/rpc/marketingcenter"
	"business/common/curl"
	"business/common/utils"
	"business/common/vars"
	"context"
	"errors"
	"time"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CampaignCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCampaignCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CampaignCreateLogic {
	return &CampaignCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CampaignCreateLogic) CampaignCreate(req *types.CampaignCreateReq) (resp *types.CampaignCreateResp, err error) {
	if err = l.svcCtx.Validator.StructCtx(l.ctx, req); err != nil {
		return nil, err
	}
	app, err := l.svcCtx.AppRpcClient.GetAppInfoByAppId(l.ctx, &appcenter.GetByAppIdReq{AppId: req.AppId})
	if err != nil {
		return nil, utils.RpcError(err, "应用信息错误")
	}
	token, err := l.svcCtx.AccountRpcClient.GetToken(l.ctx, &accountcenter.GetTokenReq{AccountId: app.AccountId})
	if err != nil {
		return nil, utils.RpcError(err, "账户信息请求错误")
	}
	if token.ExpiredAt <= time.Now().Unix() {
		token, err = l.svcCtx.AccountRpcClient.RefreshToken(l.ctx, &accountcenter.GetTokenReq{AccountId: app.AccountId})
		if err != nil {
			return nil, utils.RpcError(err, "数据异常")
		}
	}
	c, err := curl.New(l.svcCtx.Config.MarketingApis.Campaign.Create).Post().JsonData(statements.CampaignCreate{
		AdvertiserId: app.AdvertiserId,
		CampaignName: req.CampaignName,
		ProductType:  req.ProductType,
		DailyBudget:  req.DailyBudget,
		CampaignType: req.CampaignType,
		SyncFlow:     req.SyncFlow,
	})
	if err != nil {
		return nil, err
	}
	var rs statements.CampaignCreateResp
	if err = c.Request(&rs, curl.Authorization(token.TokenType+" "+token.AccessToken), curl.JsonHeader()); err != nil {
		return nil, errors.New("华为 ADS 接口调用失败：" + err.Error())
	}
	if rs.Code != "200" {
		return nil, errors.New("华为 ADS 接口调用失败：" + rs.Message)
	}
	_, err = l.svcCtx.MarketingRpcClient.CampaignCreate(l.ctx, &marketingcenter.CampaignCreateReq{
		CampaignId:   rs.Data.CampaignId,
		CampaignName: req.CampaignName,
		AdvertiserId: app.AdvertiserId,
		ProductType:  req.ProductType,
		DailyBudget:  req.DailyBudget,
		AccountId:    app.AccountId,
		AppId:        req.AppId,
		CampaignType: req.CampaignType,
		SyncFlow:     req.SyncFlow,
	})
	if err != nil {
		return nil, utils.RpcError(err, "")
	}
	return &types.CampaignCreateResp{
		Code: vars.ResponseCodeOk,
		Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		Data: nil,
	}, nil
}
