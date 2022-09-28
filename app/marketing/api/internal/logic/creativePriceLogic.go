package logic

import (
	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"
	"business/app/marketing/rpc/marketingcenter"
	"business/common/utils"
	"business/common/vars"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreativePriceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreativePriceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreativePriceLogic {
	return &CreativePriceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreativePriceLogic) CreativePrice(req *types.CreativeSizePriceReq) (resp *types.CreativeSizePriceResp, err error) {
	price, err := l.svcCtx.MarketingRpcClient.GetPositionPrice(l.ctx, &marketingcenter.PositionPriceReq{
		CreativeSizeId: req.CreativeSizeId,
		PriceType:      req.PriceType,
	})
	if err != nil {
		return nil, utils.RpcError(err, "未查询到底价信息，请联系管理员是否有拉取到")
	}
	return &types.CreativeSizePriceResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: price.Price,
	}, nil
}
