package logic

import (
	"context"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PromotionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPromotionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PromotionListLogic {
	return &PromotionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PromotionListLogic) PromotionList(req *types.PromotionListReq) (resp *types.PromotionListResp, err error) {

	return
}
