package logic

import (
	"context"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type CampaignInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCampaignInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CampaignInfoLogic {
	return &CampaignInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CampaignInfoLogic) CampaignInfo(in *marketing.CampaignInfoReq) (*marketing.PromotionInfo, error) {
	campaign, err := l.svcCtx.CampaignModel.FindOneByCampaignId(l.ctx, in.CampaignId)
	if err != nil {
		return nil, err
	}
	return &marketing.PromotionInfo{
		Id:                        campaign.Id,
		CampaignId:                campaign.CampaignId,
		CampaignName:              campaign.CampaignName,
		AccountId:                 campaign.AccountId,
		AdvertiserId:              campaign.AdvertiserId,
		OptStatus:                 campaign.OptStatus,
		CampaignDailyBudgetStatus: campaign.CampaignDailyBudgetStatus,
		ProductType:               campaign.ProductType,
		ShowStatus:                campaign.ShowStatus,
		UserBalanceStatus:         campaign.UserBalanceStatus,
		FlowResource:              campaign.FlowResource,
		SyncFlowResource:          campaign.SyncFlowResource,
		CampaignType:              campaign.CampaignType,
		TodayDailyBudget:          campaign.TodayDailyBudget,
		TomorrowDailyBudget:       campaign.TomorrowDailyBudget,
		MarketingGoal:             campaign.MarketingGoal,
		IsCallback:                campaign.IsCallback,
		CreatedAt:                 campaign.CreatedAt.Unix(),
		UpdatedAt:                 campaign.UpdatedAt.Unix(),
	}, nil
}
