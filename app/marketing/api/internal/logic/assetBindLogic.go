package logic

import (
	"business/app/marketing/rpc/marketingcenter"
	"context"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssetBindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssetBindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetBindLogic {
	return &AssetBindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssetBindLogic) AssetBind(req *types.AssetBindReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.MarketingRpcClient.BindAsset(l.ctx, &marketingcenter.AssetBindReq{
		AppId:   req.AppId,
		AssetId: req.AssetIds,
	})
	if err != nil {
		return nil, err
	}

	return successFul(), nil
}
