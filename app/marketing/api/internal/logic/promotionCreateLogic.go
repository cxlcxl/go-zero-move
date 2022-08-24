package logic

import (
	"context"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
