package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssetElementLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAssetElementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetElementLogic {
	return &AssetElementLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AssetElementLogic) AssetElement(in *marketing.AssetElementReq) (*marketing.AssetElementResp, error) {
	assets, err := l.svcCtx.AssetModel.AssetElementList(l.ctx, in.AppId, in.AssetType, in.Width, in.Height, in.FileSize, in.AssetName)
	if err != nil {
		return nil, err
	}
	var rs []*marketing.Asset
	if err = copier.Copy(&rs, assets); err != nil {
		return nil, err
	}
	return &marketing.AssetElementResp{
		Assets: rs,
	}, nil
}
