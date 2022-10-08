package logic

import (
	"business/common/vars"
	"context"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PositionCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPositionCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PositionCategoryLogic {
	return &PositionCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PositionCategoryLogic) PositionCategory() (resp *types.AdsMapResourceResp, err error) {
	return &types.AdsMapResourceResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: vars.CreativeCategory,
	}, nil
}
