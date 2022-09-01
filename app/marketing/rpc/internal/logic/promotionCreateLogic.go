package logic

import (
	"business/app/marketing/rpc/model"
	"context"
	"time"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type PromotionCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPromotionCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PromotionCreateLogic {
	return &PromotionCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PromotionCreateLogic) PromotionCreate(in *marketing.PromotionCreateReq) (*marketing.BaseResp, error) {

	_, err := l.svcCtx.CampaignModel.Insert(l.ctx, &model.Campaigns{
		CampaignId:       in.CampaignId,
		CampaignName:     in.CampaignName,
		AccountId:        in.AccountId,
		AdvertiserId:     in.AdvertiserId,
		ProductType:      in.ProductType,
		SyncFlowResource: in.SyncFlow,
		CampaignType:     in.CampaignType,
		TodayDailyBudget: in.DailyBudget,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &marketing.BaseResp{}, nil
}
