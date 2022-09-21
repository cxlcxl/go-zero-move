package logic

import (
	"business/app/application/rpc/appcenter"
	"business/app/marketing/rpc/marketing"
	"business/common/utils"
	"business/common/vars"
	"context"
	"github.com/jinzhu/copier"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CampaignInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCampaignInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CampaignInfoLogic {
	return &CampaignInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CampaignInfoLogic) CampaignInfo(req *types.CampaignInfoReq) (resp *types.CampaignInfoResp, err error) {
	info, err := l.svcCtx.MarketingRpcClient.GetCampaignInfo(l.ctx, &marketing.CampaignInfoReq{CampaignId: req.CampaignId})
	if err != nil {
		return nil, utils.RpcError(err, "计划信息请求错误")
	}
	var d types.Campaign
	if err = copier.Copy(&d, info); err != nil {
		return nil, err
	}
	app, err := l.svcCtx.AppRpcClient.GetAppInfoByAppId(l.ctx, &appcenter.GetByAppIdReq{AppId: info.AppId})
	if err != nil {
		return nil, utils.RpcError(err, "应用查询失败")
	}
	d.App = types.AppInfo{
		ProductId:           app.ProductId,
		AppName:             app.AppName,
		IconUrl:             app.IconUrl,
		AppStoreDownloadUrl: app.AppStoreDownloadUrl,
	}
	return &types.CampaignInfoResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: d,
	}, nil
}
