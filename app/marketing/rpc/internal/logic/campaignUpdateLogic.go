package logic

import (
	"context"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type CampaignUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCampaignUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CampaignUpdateLogic {
	return &CampaignUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CampaignUpdateLogic) CampaignUpdate(in *marketing.CampaignUpdateReq) (*marketing.BaseResp, error) {
	values := map[string]interface{}{
		"campaign_name":         in.CampaignName,
		"product_type":          in.ProductType,
		"today_daily_budget":    in.DailyBudget,
		"tomorrow_daily_budget": in.DailyBudget,
	}
	if err := l.svcCtx.CampaignModel.UpdateByCampaignId(l.ctx, in.CampaignId, values); err != nil {
		return nil, err
	}
	return &marketing.BaseResp{}, nil
}
