package logic

import (
	"business/app/marketing/rpc/marketingcenter"
	"business/common/utils"
	"business/common/vars"
	"context"
	"github.com/jinzhu/copier"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PromotionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPromotionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PromotionListLogic {
	return &PromotionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PromotionListLogic) PromotionList(req *types.PromotionListReq) (resp *types.PromotionListResp, err error) {
	list, err := l.svcCtx.MarketingRpcClient.PromotionList(l.ctx, &marketingcenter.PromotionListReq{
		CampaignId:   req.CampaignId,
		CampaignName: req.CampaignName,
		CampaignType: req.CampaignType,
		Page:         req.Page,
		PageSize:     req.PageSize,
	})
	if err != nil {
		return nil, utils.RpcError(err, "")
	}
	var promotions []*types.Promotion
	if err = copier.Copy(&promotions, list.Promotions); err != nil {
		return nil, err
	}
	return &types.PromotionListResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Total: list.Total,
		Data:  promotions,
		Resources: &types.ListResources{
			OptStatus:       vars.OptStatus,
			ProductType:     vars.ProductType,
			CampaignType:    vars.CampaignType,
			SyncFlow:        vars.SyncFlow,
			DailyBudgetOpts: vars.DailyBudgetOpts,
		},
	}, nil
}
