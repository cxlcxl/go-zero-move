package logic

import (
	"business/app/marketing/rpc/model"
	"context"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchInsertTrackingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchInsertTrackingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchInsertTrackingLogic {
	return &BatchInsertTrackingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchInsertTrackingLogic) BatchInsertTracking(in *marketing.BatchInsertTrackingReq) (*marketing.BaseResp, error) {
	trackings := make([]*model.Trackings, len(in.Trackings))
	for i, asset := range in.Trackings {
		trackings[i] = &model.Trackings{
			AccountId:        asset.AccountId,
			AdvertiserId:     asset.AdvertiserId,
			AppId:            asset.AppId,
			TrackingId:       asset.TrackingId,
			EffectName:       asset.EffectName,
			EffectType:       asset.EffectType,
			ClickTrackingUrl: asset.ClickTrackingUrl,
			ImpTrackingUrl:   asset.ImpTrackingUrl,
		}
	}
	err := l.svcCtx.TrackModel.BatchInsert(l.ctx, trackings, in.DisabledTrackingIds)
	if err != nil {
		return nil, err
	}
	return &marketing.BaseResp{}, nil
}
