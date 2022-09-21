package logic

import (
	"business/common/vars"
	"context"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PricingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPricingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PricingLogic {
	return &PricingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PricingLogic) Pricing() (resp *types.AdsMapResourceResp, err error) {
	return &types.AdsMapResourceResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: vars.Pricing,
	}, nil
}
