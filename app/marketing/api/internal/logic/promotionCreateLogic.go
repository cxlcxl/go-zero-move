package logic

import (
	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type PromotionCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPromotionCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PromotionCreateLogic {
	return &PromotionCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PromotionCreateLogic) PromotionCreate(req *types.PromotionCreateReq) (resp *types.PromotionCreateResp, err error) {
	if err = l.svcCtx.Validator.StructCtx(l.ctx, req); err != nil {
		return nil, err
	}
	return
}
