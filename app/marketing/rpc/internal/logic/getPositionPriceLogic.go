package logic

import (
	"context"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPositionPriceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPositionPriceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPositionPriceLogic {
	return &GetPositionPriceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPositionPriceLogic) GetPositionPrice(in *marketing.PositionPriceReq) (*marketing.PositionPriceResp, error) {
	price, err := l.svcCtx.PriceModel.FindFloorPrice(l.ctx, in.CreativeSizeId, in.PriceType)
	if err != nil {
		return nil, err
	}

	return &marketing.PositionPriceResp{
		Price: price,
	}, nil
}
