package logic

import (
	"context"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type TargetingListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTargetingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TargetingListLogic {
	return &TargetingListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TargetingListLogic) TargetingList(in *marketing.EmptyParamsReq) (*marketing.TargetingListResp, error) {
	targetings, err := l.svcCtx.TargetingModel.GetTargetings(l.ctx)
	if err != nil {
		return nil, err
	}
	var rs []*marketing.Targeting
	for _, targeting := range targetings {
		rs = append(rs, &marketing.Targeting{
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
		})
	}
	return &marketing.TargetingListResp{
		Targetings: rs,
	}, nil
}
