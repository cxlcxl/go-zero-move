package logic

import (
	"business/app/account/rpc/accountcenter"
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
	info, err := l.svcCtx.AccountRpcClient.GetAccountInfo(l.ctx, &accountcenter.AccountInfoReq{Id: req.AccountId})
	if err != nil {
		return nil, utils.RpcError(err, "账户信息请求错误")
	}
	token, err := l.svcCtx.AccountRpcClient.GetToken(l.ctx, &accountcenter.GetTokenReq{AccountId: req.AccountId})
	if err != nil {
		return nil, utils.RpcError(err, "账户信息请求错误")
	}
	if token.ExpiredAt <= time.Now().Unix() {
		return nil, errors.New(vars.ResponseMsg[vars.ResponseCodeTokenExpire])
	}
	c, err := curl.New(l.svcCtx.Config.MarketingApis.Promotion.Create).Post().JsonData(statements.CampaignCreate{
		AdvertiserId: info.AdvertiserId,
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
	_, err = l.svcCtx.MarketingRpcClient.PromotionCreate(l.ctx, &marketingcenter.PromotionCreateReq{
		CampaignId:   rs.Data.CampaignId,
		CampaignName: req.CampaignName,
		AdvertiserId: info.AdvertiserId,
		ProductType:  req.ProductType,
		DailyBudget:  req.DailyBudget,
		AccountId:    req.AccountId,
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
