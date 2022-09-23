package logic

import (
	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"
	"business/app/marketing/rpc/model"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchInsertAssetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchInsertAssetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchInsertAssetLogic {
	return &BatchInsertAssetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchInsertAssetLogic) BatchInsertAsset(in *marketing.BatchInsertAssetReq) (*marketing.BaseResp, error) {
	assets := make([]*model.Assets, len(in.Assets))
	for i, asset := range in.Assets {
		assets[i] = &model.Assets{
			AccountId:         asset.AccountId,
			AdvertiserId:      asset.AdvertiserId,
			AppId:             asset.AppId,
			AssetId:           asset.AssetId,
			AssetName:         asset.AssetName,
			AssetType:         asset.AssetType,
			FileUrl:           asset.FileUrl,
			Width:             asset.Width,
			Height:            asset.Height,
			VideoPlayDuration: asset.VideoPlayDuration,
			FileSize:          asset.FileSize,
			FileFormat:        asset.FileFormat,
			FileHashSha256:    asset.FileHashSha256,
		}
	}
	err := l.svcCtx.AssetModel.BatchInsert(l.ctx, assets)
	if err != nil {
		return nil, err
	}
	return &marketing.BaseResp{}, nil
}
