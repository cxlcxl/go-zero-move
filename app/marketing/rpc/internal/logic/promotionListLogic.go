package logic

import (
	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"
	"business/common/utils"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type PromotionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPromotionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PromotionListLogic {
	return &PromotionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PromotionListLogic) PromotionList(in *marketing.PromotionListReq) (*marketing.PromotionListResp, error) {
	offset, limit := utils.Pagination(in.Page, in.PageSize)
	list, total, err := l.svcCtx.CampaignModel.CampaignList(l.ctx, in.CampaignId, in.CampaignName, in.CampaignType, offset, limit)
	if err != nil {
		return nil, err
	}
	promotions := make([]*marketing.PromotionInfo, len(list))
	for i, campaigns := range list {
		promotions[i] = &marketing.PromotionInfo{
			Id:                        campaigns.Id,
			CampaignId:                campaigns.CampaignId,
			CampaignName:              campaigns.CampaignName,
			AccountId:                 campaigns.AccountId,
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
	return &marketing.PromotionListResp{
		Promotions: promotions,
		Total:      total,
	}, nil
}
