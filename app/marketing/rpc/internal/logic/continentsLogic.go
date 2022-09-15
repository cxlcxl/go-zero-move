package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContinentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewContinentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContinentsLogic {
	return &ContinentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ContinentsLogic) Continents(in *marketing.EmptyParamsReq) (*marketing.ContinentResp, error) {
	continents, err := l.svcCtx.ContinentModel.Continents(l.ctx)
	if err != nil {
		return nil, err
	}
	var d []*marketing.ContinentResp_Continent
	if err = copier.Copy(&d, continents); err != nil {
		return nil, err
	}
	return &marketing.ContinentResp{
		Continents: d,
	}, nil
}
