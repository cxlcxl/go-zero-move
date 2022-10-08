package logic

import (
	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"
	"context"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPositionElementLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPositionElementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPositionElementLogic {
	return &GetPositionElementLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPositionElementLogic) GetPositionElement(in *marketing.PositionElementReq) (*marketing.PositionElementResp, error) {
	elements, err := l.svcCtx.ElementModel.PositionElement(l.ctx, in.CreativeSizeId, in.SubType, in.Width, in.Height)
	if err != nil {
		return nil, err
	}
	var __elements []*marketing.PositionElementResp_Element
	if err = copier.Copy(&__elements, elements); err != nil {
		return nil, err
	}
	return &marketing.PositionElementResp{
		Elements: __elements,
	}, nil
}
