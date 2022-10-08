package logic

import (
	"context"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type CampaignBindAppLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCampaignBindAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CampaignBindAppLogic {
	return &CampaignBindAppLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CampaignBindAppLogic) CampaignBindApp(in *marketing.CampaignBindAppReq) (*marketing.BaseResp, error) {
	if err := l.svcCtx.CampaignModel.BindApp(l.ctx, in.CampaignId, in.AppId); err != nil {
		return nil, err
	}

	return &marketing.BaseResp{}, nil
}
