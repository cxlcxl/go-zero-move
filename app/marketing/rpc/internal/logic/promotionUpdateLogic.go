package logic

import (
	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type PromotionUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPromotionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PromotionUpdateLogic {
	return &PromotionUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PromotionUpdateLogic) PromotionUpdate(in *marketing.PromotionUpdateReq) (*marketing.BaseResp, error) {
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
