package logic

import (
	"context"

	"business/app/application/rpc/apps"
	"business/app/application/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppsByAppNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAppsByAppNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppsByAppNameLogic {
	return &GetAppsByAppNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAppsByAppNameLogic) GetAppsByAppName(in *apps.GetByAppNameReq) (*apps.AppsResp, error) {
	appInfos, err := l.svcCtx.AppModel.GetAppsByAppName(l.ctx, in.AppName)
	if err != nil {
		return nil, err
	}
	var info []*apps.AppInfo
	for _, a := range appInfos {
		info = append(info, &apps.AppInfo{
			Id:                  a.Id,
			AppName:             a.AppName,
			AppId:               a.AppId,
			AccountId:           a.AccountId,
			PkgName:             a.PkgName,
			AdvertiserId:        a.AdvertiserId,
			ProductId:           a.ProductId,
			AppStoreDownloadUrl: a.AppStoreDownloadUrl,
			IconUrl:             a.IconUrl,
			Channel:             a.Channel,
			AppType:             a.AppType,
			Tags:                a.Tags,
			CreatedAt:           a.CreatedAt.Unix(),
			UpdatedAt:           a.UpdatedAt.Unix(),
		})
	}
	return &apps.AppsResp{
		Apps: info,
	}, nil
}
