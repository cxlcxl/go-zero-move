package logic

import (
	"business/app/marketing/rpc/marketingcenter"
	"business/common/utils"
	"business/common/vars"
	"context"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PositionPriceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPositionPriceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PositionPriceLogic {
	return &PositionPriceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PositionPriceLogic) PositionPrice(req *types.PositionPriceReq) (resp *types.PositionPriceResp, err error) {
	price, err := l.svcCtx.MarketingRpcClient.GetPositionPrice(l.ctx, &marketingcenter.PositionPriceReq{
		CreativeSizeId: req.CreativeSizeId,
		PriceType:      req.PriceType,
	})
	if err != nil {
		return nil, utils.RpcError(err, "未查询到底价信息，请联系管理员是否有拉取到")
	}
	return &types.PositionPriceResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: price.Price,
	}, nil
}
