package logic

import (
	"context"

	"business/app/application/rpc/apps"
	"business/app/application/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAppInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppInfoLogic {
	return &GetAppInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAppInfoLogic) GetAppInfo(in *apps.AppInfoReq) (*apps.AppInfo, error) {
	app, err := l.svcCtx.AppModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &apps.AppInfo{
		Id:                  app.Id,
		AppName:             app.AppName,
		AppId:               app.AppId,
		AdvertiserId:        app.AdvertiserId,
		IconUrl:             app.IconUrl,
		AccountId:           app.AccountId,
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
