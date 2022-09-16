package logic

import (
	"context"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTargetingByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTargetingByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTargetingByNameLogic {
	return &GetTargetingByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTargetingByNameLogic) GetTargetingByName(in *marketing.GetTargetingByNameReq) (*marketing.Targeting, error) {
	targeting, err := l.svcCtx.TargetingModel.FindOneByTargetingName(l.ctx, in.TargetingName)
	if err != nil {
		return nil, err
	}
	return &marketing.Targeting{
		Id:                 targeting.Id,
		AccountId:          targeting.AccountId,
		AdvertiserId:       targeting.AdvertiserId,
		TargetingId:        targeting.TargetingId,
		TargetingName:      targeting.TargetingName,
		TargetingType:      targeting.TargetingType,
		LocationType:       targeting.LocationType,
		IncludeLocation:    targeting.IncludeLocation,
		ExcludeLocation:    targeting.ExcludeLocation,
		Carriers:           targeting.Carriers,
		Language:           targeting.Language,
		Age:                targeting.Age,
		Gender:             targeting.Gender,
		AppCategory:        targeting.AppCategory,
		AppCategories:      targeting.AppCategories,
		InstalledApps:      targeting.InstalledApps,
		AppInterest:        targeting.AppInterest,
		AppInterests:       targeting.AppInterests,
		Series:             targeting.Series,
		NetworkType:        targeting.NetworkType,
		NotAudiences:       targeting.NotAudiences,
		Audiences:          targeting.Audiences,
		AppCategoryOfMedia: targeting.AppCategoryOfMedia,
		CreatedAt:          targeting.CreatedAt.Unix(),
		UpdatedAt:          targeting.UpdatedAt.Unix(),
	}, nil
}
