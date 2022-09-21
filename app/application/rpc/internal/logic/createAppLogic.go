package logic

import (
	"business/app/application/rpc/model"
	"context"
	"time"

	"business/app/application/rpc/apps"
	"business/app/application/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAppLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAppLogic {
	return &CreateAppLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateAppLogic) CreateApp(in *apps.CreateAppReq) (*apps.BaseResp, error) {
	_, err := l.svcCtx.AppModel.Insert(l.ctx, &model.Apps{
		AccountId:    in.AccountId,
		AdvertiserId: in.AdvertiserId,
		AppId:        in.AppId,
		AppName:      in.AppName,
		PkgName:      in.PkgName,
		Channel:      in.Channel,
		AppType:      in.AppType,
		Tags:         in.Tags,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &apps.BaseResp{}, nil
}
