package logic

import (
	"business/app/application/rpc/appcenter"
	"business/common/utils"
	"business/common/vars"
	"context"
	"github.com/jinzhu/copier"

	"business/app/application/api/internal/svc"
	"business/app/application/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CampaignAppListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCampaignAppListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CampaignAppListLogic {
	return &CampaignAppListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CampaignAppListLogic) CampaignAppList(req *types.CampaignAppReq) (resp *types.CampaignAppListResp, err error) {
	list, err := l.svcCtx.AppRpcClient.AppList(l.ctx, &appcenter.AppListReq{
		Page: req.Page, PageSize: req.PageSize, AppName: req.AppName,
	})
	if err != nil {
		return nil, utils.RpcError(err, "没有相关应用信息")
	}
	var rs []*types.CampaignAppInfo
	if err = copier.Copy(&rs, list.Apps); err != nil {
		return nil, err
	}
	return &types.CampaignAppListResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Total: list.Total,
		Data:  rs,
	}, nil
}
