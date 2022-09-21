package logic

import (
	"business/common/utils"
	"context"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type CampaignListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCampaignListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CampaignListLogic {
	return &CampaignListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CampaignListLogic) CampaignList(in *marketing.CampaignListReq) (*marketing.CampaignListResp, error) {
	offset, limit := utils.Pagination(in.Page, in.PageSize)
	list, total, err := l.svcCtx.CampaignModel.CampaignList(l.ctx, in.AppId, in.CampaignId, in.CampaignName, in.CampaignType, offset, limit)
	if err != nil {
		return nil, err
	}
	promotions := make([]*marketing.CampaignInfo, len(list))
	for i, campaigns := range list {
		promotions[i] = &marketing.CampaignInfo{
			Id:                        campaigns.Id,
			CampaignId:                campaigns.CampaignId,
			CampaignName:              campaigns.CampaignName,
			AccountId:                 campaigns.AccountId,
			AppId:                     campaigns.AppId,
			AdvertiserId:              campaigns.AdvertiserId,
			OptStatus:                 campaigns.OptStatus,
			CampaignDailyBudgetStatus: campaigns.CampaignDailyBudgetStatus,
			ProductType:               campaigns.ProductType,
			ShowStatus:                campaigns.ShowStatus,
			UserBalanceStatus:         campaigns.UserBalanceStatus,
			FlowResource:              campaigns.FlowResource,
			SyncFlowResource:          campaigns.SyncFlowResource,
			CampaignType:              campaigns.CampaignType,
			TodayDailyBudget:          campaigns.TodayDailyBudget,
			TomorrowDailyBudget:       campaigns.TomorrowDailyBudget,
			MarketingGoal:             campaigns.MarketingGoal,
			IsCallback:                campaigns.IsCallback,
			CreatedAt:                 campaigns.CreatedAt.Unix(),
			UpdatedAt:                 campaigns.UpdatedAt.Unix(),
		}
	}
	return &marketing.CampaignListResp{
		Campaigns: promotions,
		Total:     total,
	}, nil
}
