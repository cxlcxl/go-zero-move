package logic

import (
	"context"

	"business/app/ads/api/internal/svc"
	"business/app/ads/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
