package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCountriesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCountriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCountriesLogic {
	return &GetCountriesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCountriesLogic) GetCountries(in *marketing.EmptyParamsReq) (*marketing.CountriesResp, error) {
	countries, err := l.svcCtx.RegionModel.GetCountries(l.ctx)
	if err != nil {
		return nil, err
	}
	var d []*marketing.CountriesResp_OverseasRegionCountry
	if err = copier.Copy(&d, countries); err != nil {
		return nil, err
	}
	return &marketing.CountriesResp{
		Countries: d,
	}, nil
}
