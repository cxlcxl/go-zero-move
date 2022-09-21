package logic

import (
	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"
	"business/app/marketing/rpc/marketing"
	"business/common/utils"
	"business/common/vars"
	"context"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreativeQueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreativeQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreativeQueryLogic {
	return &CreativeQueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreativeQueryLogic) CreativeQuery(req *types.CreativeQueryReq) (resp *types.CreativeQueryResp, err error) {
	position, err := l.svcCtx.MarketingRpcClient.GetPositions(l.ctx, &marketing.PositionListReq{
		Category:  req.Category,
		AccountId: req.AccountId,
	})
	if err != nil {
		return nil, utils.RpcError(err, "版位数据请求异常")
	}
	var rs []*types.CreativeSizeInfo
	for _, info := range position.CreativeSizeInfoList {
		var samples []*types.SampleInfo
		var placements []*types.PlacementInfo
		if err = copier.Copy(&samples, info.CreativeSizeSampleList); err != nil {
			return nil, err
		}
		if err = copier.Copy(&placements, info.CreativeSizePlacementList); err != nil {
			return nil, err
		}
		rs = append(rs, &types.CreativeSizeInfo{
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
	return &types.CreativeQueryResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: rs,
	}, nil
}
