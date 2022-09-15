package logic

import (
	"business/app/marketing/rpc/marketingcenter"
	"business/common/utils"
	"business/common/vars"
	"context"
	"github.com/jinzhu/copier"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TargetingLocationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTargetingLocationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TargetingLocationLogic {
	return &TargetingLocationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TargetingLocationLogic) TargetingLocation() (resp *types.LocationResp, err error) {
	continents, err := l.svcCtx.MarketingRpcClient.Continents(l.ctx, &marketingcenter.EmptyParamsReq{})
	if err != nil {
		return nil, utils.RpcError(err, "大洲信息未填充")
	}
	countries, err := l.svcCtx.MarketingRpcClient.GetCountries(l.ctx, &marketingcenter.EmptyParamsReq{})
	if err != nil {
		return nil, utils.RpcError(err, "国家信息为空")
	}
	var rsContinents []*types.Continent
	if err = copier.Copy(&rsContinents, continents.Continents); err != nil {
		return nil, err
	}
	var rsCountries []*types.Country
	if err = copier.Copy(&rsCountries, countries.Countries); err != nil {
		return nil, err
	}
	return &types.LocationResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: types.ContinentCountry{
			Continents: rsContinents,
			Countries:  rsCountries,
		},
	}, nil
}
