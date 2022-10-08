package logic

import (
	"business/app/marketing/rpc/marketingcenter"
	"business/common/utils"
	"context"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CampaignBindAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCampaignBindAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CampaignBindAppLogic {
	return &CampaignBindAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CampaignBindAppLogic) CampaignBindApp(req *types.CampaignBindAppReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.MarketingRpcClient.CampaignBindApp(l.ctx, &marketingcenter.CampaignBindAppReq{CampaignId: req.CampaignIds, AppId: req.AppId})
	if err != nil {
		return nil, utils.RpcError(err, "")
	}
	return successFul(), nil
}
