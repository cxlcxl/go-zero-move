package logic

import (
	"business/app/marketing/rpc/model"
	"context"
	"time"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type TargetingCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTargetingCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TargetingCreateLogic {
	return &TargetingCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TargetingCreateLogic) TargetingCreate(in *marketing.Targeting) (*marketing.BaseResp, error) {
	_, err := l.svcCtx.TargetingModel.Insert(l.ctx, &model.Targetings{
		AccountId:          in.AccountId,
		AdvertiserId:       in.AdvertiserId,
		TargetingId:        in.TargetingId,
		TargetingName:      in.TargetingName,
		TargetingType:      in.TargetingType,
		LocationType:       in.LocationType,
		IncludeLocation:    in.IncludeLocation,
		ExcludeLocation:    in.ExcludeLocation,
		Carriers:           in.Carriers,
		Language:           in.Language,
		Age:                in.Age,
		Gender:             in.Gender,
		AppCategory:        in.AppCategory,
		AppCategories:      in.AppCategories,
		InstalledApps:      in.InstalledApps,
		AppInterest:        in.AppInterest,
		AppInterests:       in.AppInterests,
		Series:             in.Series,
		NetworkType:        in.NetworkType,
		NotAudiences:       in.NotAudiences,
		Audiences:          in.Audiences,
		AppCategoryOfMedia: in.AppCategoryOfMedia,
		CreatedAt:          time.Unix(in.CreatedAt, 0),
		UpdatedAt:          time.Unix(in.CreatedAt, 0),
	})
	if err != nil {
		return nil, err
	}
	return &marketing.BaseResp{}, nil
}
