package logic

import (
	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPositionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPositionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPositionsLogic {
	return &GetPositionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPositionsLogic) GetPositions(in *marketing.PositionListReq) (*marketing.PositionListResp, error) {
	positions, err := l.svcCtx.PositionModel.PositionList(l.ctx, in.Category, in.ProductType, in.AccountId)
	if err != nil {
		return nil, err
	}
	var (
		creativeSizeIds  []string
		creativeSizeList []*marketing.CreativeSizeInfo
	)
	for _, position := range positions {
		creativeSizeIds = append(creativeSizeIds, position.CreativeSizeId)
		creativeSizeList = append(creativeSizeList, &marketing.CreativeSizeInfo{
			CreativeSizeId:             position.CreativeSizeId,
			CreativeSizeNameDsp:        position.CreativeSizeNameDsp,
			CreativeSizeDescription:    position.CreativeSizeDescription,
			SupportProductType:         position.SupportProductType,
			SupportObjectiveType:       position.SupportObjectiveType,
			IsSupportTimePeriod:        position.IsSupportTimePeriod,
			IsSupportMultipleCreatives: position.IsSupportMultipleCreatives,
			SupportPriceType:           position.SupportPriceType,
		})
	}
	samples, err := l.svcCtx.SampleModel.SampleList(l.ctx, creativeSizeIds)
	if err != nil {
		return nil, err
	}
	placements, err := l.svcCtx.PlacementModel.PlacementList(l.ctx, creativeSizeIds)
	if err != nil {
		return nil, err
	}
	sampleTmp := make(map[string][]*marketing.CreativeSizeSample)
	for _, sample := range samples {
		sampleTmp[sample.CreativeSizeId] = append(sampleTmp[sample.CreativeSizeId], &marketing.CreativeSizeSample{
			CreativeSizeSample: sample.CreativeSizeSample,
			PreviewTitle:       sample.PreviewTitle,
		})
	}
	placementTmp := make(map[string][]*marketing.CreativeSizePlacement)
	for _, placement := range placements {
		placementTmp[placement.CreativeSizeId] = append(placementTmp[placement.CreativeSizeId], &marketing.CreativeSizePlacement{
			PlacementSizeId:     placement.PlacementSizeId,
			CreativeSize:        placement.CreativeSize,
			CreativeSizeSubType: placement.CreativeSizeSubType,
		})
	}
	for i, creativeSize := range creativeSizeList {
		if v, ok := sampleTmp[creativeSize.CreativeSizeId]; ok {
			creativeSizeList[i].CreativeSizeSampleList = v
		}
		if v, ok := placementTmp[creativeSize.CreativeSizeId]; ok {
			creativeSizeList[i].CreativeSizePlacementList = v
		}
	}
	return &marketing.PositionListResp{
		CreativeSizeInfoList: creativeSizeList,
	}, nil
}
