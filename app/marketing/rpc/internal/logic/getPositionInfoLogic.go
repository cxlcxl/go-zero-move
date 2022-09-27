package logic

import (
	"context"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPositionInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPositionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPositionInfoLogic {
	return &GetPositionInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPositionInfoLogic) GetPositionInfo(in *marketing.CreativeSizeInfoReq) (*marketing.CreativeSizeInfoResp, error) {
	creativeSize, err := l.svcCtx.PositionModel.FindOneByCreativeSizeId(l.ctx, in.CreativeSizeId)
	if err != nil {
		return nil, err
	}

	return &marketing.CreativeSizeInfoResp{
		CreativeSizeInfo: &marketing.CreativeSizeInfo{
			CreativeSizeId:             creativeSize.CreativeSizeId,
			CreativeSizeNameDsp:        creativeSize.CreativeSizeNameDsp,
			CreativeSizeDescription:    creativeSize.CreativeSizeDescription,
			SupportProductType:         creativeSize.SupportProductType,
			SupportObjectiveType:       creativeSize.SupportObjectiveType,
			IsSupportTimePeriod:        creativeSize.IsSupportTimePeriod,
			IsSupportMultipleCreatives: creativeSize.IsSupportMultipleCreatives,
			SupportPriceType:           creativeSize.SupportPriceType,
		},
		AccountId:    creativeSize.AccountId,
		AdvertiserId: creativeSize.AdvertiserId,
	}, nil
}
