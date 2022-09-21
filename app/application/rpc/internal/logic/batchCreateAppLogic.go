package logic

import (
	"business/app/application/rpc/model"
	"context"
	"time"

	"business/app/application/rpc/apps"
	"business/app/application/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchCreateAppLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchCreateAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchCreateAppLogic {
	return &BatchCreateAppLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchCreateAppLogic) BatchCreateApp(in *apps.BatchCreateAppReq) (*apps.BaseResp, error) {
	_apps := make([]*model.Apps, len(in.Apps))
	now := time.Now()
	for i, app := range in.Apps {
		_apps[i] = &model.Apps{
			AccountId:           app.AccountId,
			AdvertiserId:        app.AdvertiserId,
			AppId:               app.AppId,
			AppName:             app.AppName,
			AppType:             app.AppType,
			PkgName:             app.PkgName,
			IconUrl:             app.IconUrl,
			ProductId:           app.ProductId,
			AppStoreDownloadUrl: app.AppStoreDownloadUrl,
			CreatedAt:           now,
			UpdatedAt:           now,
		}
	}
	err := l.svcCtx.AppModel.BatchInsert(l.ctx, _apps)
	if err != nil {
		return nil, err
	}
	return &apps.BaseResp{}, nil
}
