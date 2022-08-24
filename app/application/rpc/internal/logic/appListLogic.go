package logic

import (
	"business/app/application/rpc/apps"
	"business/app/application/rpc/internal/svc"
	"business/common/utils"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppListLogic {
	return &AppListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AppListLogic) AppList(in *apps.AppListReq) (*apps.AppListResp, error) {
	offset, limit := utils.Pagination(in.Page, in.PageSize)
	list, total, err := l.svcCtx.AppModel.AppList(l.ctx, in.AppId, in.AppName, in.AppType, in.Channel, offset, limit)
	if err != nil {
		return nil, err
	}
	var info []*apps.AppInfo
	if total > 0 {
		for _, app := range list {
			info = append(info, &apps.AppInfo{
				Id:        app.Id,
				AppName:   app.AppName,
				AppId:     app.AppId,
				AccountId: app.AccountId,
				PkgName:   app.PkgName,
				Channel:   app.Channel,
				AppType:   app.AppType,
				Tags:      app.Tags,
				CreatedAt: app.CreatedAt.Unix(),
				UpdatedAt: app.UpdatedAt.Unix(),
			})
		}
	}
	return &apps.AppListResp{
		Total: total,
		Apps:  info,
	}, nil
}
