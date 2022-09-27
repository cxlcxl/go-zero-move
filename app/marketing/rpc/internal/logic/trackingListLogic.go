package logic

import (
	"context"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type TrackingListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTrackingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TrackingListLogic {
	return &TrackingListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TrackingListLogic) TrackingList(in *marketing.TrackingListReq) (*marketing.TrackingListResp, error) {
	list, err := l.svcCtx.TrackModel.TrackingList(l.ctx, in.AppId)
	if err != nil {
		return nil, err
	}
	trackings := make([]*marketing.TrackingListResp_Tracking, len(list))
	for i, t := range list {
		trackings[i] = &marketing.TrackingListResp_Tracking{
			AppId:        t.AppId,
			AccountId:    t.AccountId,
			AdvertiserId: t.AdvertiserId,
			TrackingId:   t.TrackingId,
			EffectType:   t.EffectType,
			EffectName:   t.EffectName,
		}
	}
	return &marketing.TrackingListResp{
		Trackings: trackings,
	}, nil
}
