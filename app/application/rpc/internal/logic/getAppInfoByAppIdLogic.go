package logic

import (
	"context"

	"business/app/application/rpc/apps"
	"business/app/application/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppInfoByAppIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAppInfoByAppIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppInfoByAppIdLogic {
	return &GetAppInfoByAppIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAppInfoByAppIdLogic) GetAppInfoByAppId(in *apps.GetByAppIdReq) (*apps.AppInfo, error) {
	app, err := l.svcCtx.AppModel.FindOneByAppId(l.ctx, in.AppId)
	if err != nil {
		return nil, err
	}
	return &apps.AppInfo{
		Id:                  app.Id,
		AppName:             app.AppName,
		AppId:               app.AppId,
		AccountId:           app.AccountId,
		AdvertiserId:        app.AdvertiserId,
		IconUrl:             app.IconUrl,
		ProductId:           app.ProductId,
		AppStoreDownloadUrl: app.AppStoreDownloadUrl,
		PkgName:             app.PkgName,
		Channel:             app.Channel,
		AppType:             app.AppType,
		Tags:                app.Tags,
		CreatedAt:           app.CreatedAt.Unix(),
		UpdatedAt:           app.UpdatedAt.Unix(),
	}, nil
}
