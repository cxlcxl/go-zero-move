package logic

import (
	"business/app/marketing/rpc/model"
	"context"
	"time"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type CampaignCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCampaignCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CampaignCreateLogic {
	return &CampaignCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CampaignCreateLogic) CampaignCreate(in *marketing.CampaignCreateReq) (*marketing.BaseResp, error) {
	_, err := l.svcCtx.CampaignModel.Insert(l.ctx, &model.Campaigns{
		AppId:            in.AppId,
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
