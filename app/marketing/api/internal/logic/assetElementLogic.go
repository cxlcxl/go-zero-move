package logic

import (
	"business/app/marketing/rpc/marketingcenter"
	"business/common/utils"
	"business/common/vars"
	"context"
	"github.com/jinzhu/copier"
	"strings"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssetElementLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssetElementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetElementLogic {
	return &AssetElementLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssetElementLogic) AssetElement(req *types.AssetElementReq) (resp *types.AssetElementResp, err error) {
	if err = l.svcCtx.Validator.StructCtx(l.ctx, req); err != nil {
		return nil, err
	}
	assetType := ""
	if strings.ToLower(req.FileType) == "video" {
		assetType = vars.AssetTypeVideo
	} else {
		assetType = vars.AssetTypePicture
	}
	elements, err := l.svcCtx.MarketingRpcClient.AssetElement(l.ctx, &marketingcenter.AssetElementReq{
		FileSize:  req.FileSizeKbLimit * 1024,
		AssetName: req.AssetName,
		AssetType: assetType,
		Width:     req.Width,
		Height:    req.Height,
		AppId:     req.AppId,
	})
	if err != nil {
		return nil, utils.RpcError(err, "没有相关素材")
	}
	var rs []*types.AssetElement
	if err = copier.Copy(&rs, elements.Assets); err != nil {
		return nil, err
	}
	return &types.AssetElementResp{
		BaseResp: successFul(),
		Data:     rs,
	}, nil
}
