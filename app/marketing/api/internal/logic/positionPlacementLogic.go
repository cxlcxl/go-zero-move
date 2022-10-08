package logic

import (
	"business/app/marketing/rpc/marketingcenter"
	"business/common/utils"
	"business/common/vars"
	"context"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PositionPlacementLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPositionPlacementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PositionPlacementLogic {
	return &PositionPlacementLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PositionPlacementLogic) PositionPlacement(req *types.PositionPlacementReq) (resp *types.PositionPlacementResp, err error) {
	placements, err := l.svcCtx.MarketingRpcClient.GetPositionPlacement(l.ctx, &marketingcenter.PositionPlacementReq{CreativeSizeId: req.CreativeSizeId})
	if err != nil {
		return nil, utils.RpcError(err, "没有查到此版位对应的创意子类型")
	}
	rs := make(map[string][]string)
	for _, placement := range placements.Placements {
		rs[placement.SubType] = append(rs[placement.SubType], placement.CreativeSize)
	}
	return &types.PositionPlacementResp{
		BaseResp: successFul(),
		Data: types.PlacementResp{
			Placements: rs,
			SubTypes:   vars.PositionSubType,
		},
	}, nil
}
