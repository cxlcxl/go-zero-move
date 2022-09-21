package logic

import (
	"business/app/application/rpc/model"
	"context"
	"time"

	"business/app/application/rpc/apps"
	"business/app/application/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAppLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAppLogic {
	return &UpdateAppLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAppLogic) UpdateApp(in *apps.AppInfo) (*apps.BaseResp, error) {
	err := l.svcCtx.AppModel.Update(l.ctx, &model.Apps{
		Id:           in.Id,
		AccountId:    in.AccountId,
		AdvertiserId: in.AdvertiserId,
		AppId:        in.AppId,
		AppName:      in.AppName,
		PkgName:      in.PkgName,
		Channel:      in.Channel,
		AppType:      in.AppType,
		Tags:         in.Tags,
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &apps.BaseResp{}, nil
}
