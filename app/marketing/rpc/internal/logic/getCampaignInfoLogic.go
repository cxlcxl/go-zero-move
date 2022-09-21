package logic

import (
	"context"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCampaignInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCampaignInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCampaignInfoLogic {
	return &GetCampaignInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCampaignInfoLogic) GetCampaignInfo(in *marketing.CampaignInfoReq) (*marketing.CampaignInfo, error) {
	campaign, err := l.svcCtx.CampaignModel.FindOneByCampaignId(l.ctx, in.CampaignId)
	if err != nil {
		return nil, err
	}
	return &marketing.CampaignInfo{
		Id:                        campaign.Id,
		CampaignId:                campaign.CampaignId,
		CampaignName:              campaign.CampaignName,
		AccountId:                 campaign.AccountId,
		AppId:                     campaign.AppId,
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
