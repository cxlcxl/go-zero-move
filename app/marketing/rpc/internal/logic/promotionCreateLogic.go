package logic

import (
	"context"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type PromotionCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPromotionCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PromotionCreateLogic {
	return &PromotionCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PromotionCreateLogic) PromotionCreate(in *marketing.PromotionCreateReq) (*marketing.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &marketing.BaseResp{}, nil
}
