package logic

import (
	"business/app/marketing/rpc/marketing"
	"business/common/utils"
	"business/common/vars"
	"context"
	"github.com/jinzhu/copier"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PositionQueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPositionQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PositionQueryLogic {
	return &PositionQueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PositionQueryLogic) PositionQuery(req *types.PositionQueryReq) (resp *types.PositionQueryResp, err error) {
	position, err := l.svcCtx.MarketingRpcClient.GetPositions(l.ctx, &marketing.PositionListReq{
		Category:    req.Category,
		AccountId:   req.AccountId,
		ProductType: req.ProductType,
	})
	if err != nil {
		return nil, utils.RpcError(err, "版位数据请求异常")
	}
	var rs []*types.PositionInfo
	for _, info := range position.CreativeSizeInfoList {
		var samples []*types.SampleInfo
		var placements []*types.PlacementInfo
		if err = copier.Copy(&samples, info.CreativeSizeSampleList); err != nil {
			return nil, err
		}
		if err = copier.Copy(&placements, info.CreativeSizePlacementList); err != nil {
			return nil, err
		}
		rs = append(rs, &types.PositionInfo{
			CreativeSizeId:              info.CreativeSizeId,
			CreativeSizeNameDsp:         info.CreativeSizeNameDsp,
			CreativeSizeNameDescription: info.CreativeSizeDescription,
			SupportProductType:          info.SupportProductType,
			IsSupportTimePeriod:         info.IsSupportTimePeriod,
			IsSupportMultipleCreatives:  info.IsSupportMultipleCreatives,
			SupportPriceType:            info.SupportPriceType,
			Samples:                     samples,
			Placements:                  placements,
		})
	}
	return &types.PositionQueryResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: rs,
	}, nil
}
