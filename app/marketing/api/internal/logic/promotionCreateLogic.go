package logic

import (
	"business/app/account/rpc/accountcenter"
	"business/app/marketing/api/internal/statements"
	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"
	"business/app/marketing/rpc/marketingcenter"
	"business/common/curl"
	"business/common/utils"
	"business/common/vars"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type PromotionCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPromotionCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PromotionCreateLogic {
	return &PromotionCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PromotionCreateLogic) PromotionCreate(req *types.PromotionCreateReq) (resp *types.PromotionCreateResp, err error) {
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
	c, err := curl.New(l.svcCtx.Config.MarketingApis.Promotion.Create).Post().JsonData(statements.PromotionCreate{
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
	var rs statements.PromotionCreateResp
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
		return nil, err
	}
	return &types.PromotionCreateResp{
		Code: vars.ResponseCodeOk,
		Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		Data: nil,
	}, nil
}
