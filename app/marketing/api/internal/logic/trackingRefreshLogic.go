package logic

import (
	"business/common/vars"
	"context"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TrackingRefreshLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTrackingRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TrackingRefreshLogic {
	return &TrackingRefreshLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TrackingRefreshLogic) TrackingRefresh(req *types.TrackingListReq) (resp *types.TrackingListResp, err error) {
	trackingListLogic := NewTrackingListLogic(l.ctx, l.svcCtx)
	if err = trackingListLogic.pullTracking(req.AppId); err != nil {
		return nil, err
	}
	return &types.TrackingListResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: trackingListLogic.trackings,
	}, nil
}
