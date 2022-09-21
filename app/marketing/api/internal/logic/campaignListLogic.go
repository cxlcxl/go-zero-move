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

type CampaignListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCampaignListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CampaignListLogic {
	return &CampaignListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CampaignListLogic) CampaignList(req *types.CampaignListReq) (resp *types.CampaignListResp, err error) {
	list, err := l.svcCtx.MarketingRpcClient.CampaignList(l.ctx, &marketingcenter.CampaignListReq{
		AppId:        req.AppId,
		CampaignId:   req.CampaignId,
		CampaignName: req.CampaignName,
		CampaignType: req.CampaignType,
		Page:         req.Page,
		PageSize:     req.PageSize,
	})
	if err != nil {
		return nil, utils.RpcError(err, "")
	}
	var campaigns []*types.Campaign
	if err = copier.Copy(&campaigns, list.Campaigns); err != nil {
		return nil, err
	}
	return &types.CampaignListResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: types.CampaignListData{
			Campaigns: campaigns,
			Total:     list.Total,
		},
	}, nil
}
