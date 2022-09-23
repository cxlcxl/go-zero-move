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

type AssetListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetListLogic {
	return &AssetListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssetListLogic) AssetList(req *types.AssetListReq) (resp *types.AssetListResp, err error) {
	list, err := l.svcCtx.MarketingRpcClient.AssetList(l.ctx, &marketingcenter.AssetListReq{
		Page:      req.Page,
		PageSize:  req.PageSize,
		AssetType: req.AssetType,
		Width:     req.Width,
		Height:    req.Height,
		AppId:     req.AppId,
	})
	if err != nil {
		return nil, utils.RpcError(err, "")
	}
	var rs []*types.Asset
	if err = copier.Copy(&rs, list.Assets); err != nil {
		return nil, err
	}
	return &types.AssetListResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: types.AssetListData{
			Total:     list.Total,
			List:      rs,
			AssetType: vars.AssetType,
		},
	}, nil
}
