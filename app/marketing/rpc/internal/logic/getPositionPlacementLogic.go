package logic

import (
	"context"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPositionPlacementLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPositionPlacementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPositionPlacementLogic {
	return &GetPositionPlacementLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPositionPlacementLogic) GetPositionPlacement(in *marketing.PositionPlacementReq) (*marketing.PositionPlacementResp, error) {
	placements, err := l.svcCtx.PlacementModel.PlacementsByCreativeSizeId(l.ctx, in.CreativeSizeId)
	if err != nil {
		return nil, err
	}
	rs := make([]*marketing.PositionPlacementResp_Placement, len(placements))
	for i, placement := range placements {
		rs[i] = &marketing.PositionPlacementResp_Placement{
			CreativeSize: placement.CreativeSize,
			SubType:      placement.CreativeSizeSubType,
		}
	}
	return &marketing.PositionPlacementResp{
		Placements: rs,
	}, nil
}
