package logic

import (
	"business/common/vars"
	"context"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CampaignResourcesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCampaignResourcesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CampaignResourcesLogic {
	return &CampaignResourcesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CampaignResourcesLogic) CampaignResources() (resp *types.CampaignResources, err error) {
	return &types.CampaignResources{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: &types.Resources{
			OptStatus:       vars.OptStatus,
			ProductType:     vars.ProductType,
			CampaignType:    vars.CampaignType,
			SyncFlow:        vars.SyncFlow,
			DailyBudgetOpts: vars.DailyBudgetOpts,
			ShowStatus:      vars.CampaignShowStatus,
			CampaignDaily:   vars.CampaignDaily,
			FlowResource:    vars.FlowResource,
			BalanceStatus:   vars.BalanceStatus,
		},
	}, nil
}
