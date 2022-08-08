package logic

import (
	"context"

	"business/app/application/rpc/apps"
	"business/app/application/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppsByAppIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAppsByAppIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppsByAppIdsLogic {
	return &GetAppsByAppIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAppsByAppIdsLogic) GetAppsByAppIds(in *apps.GetByAppIdsReq) (*apps.AppsResp, error) {
	appInfos, err := l.svcCtx.AppModel.GetAppsByAppIds(l.ctx, in.AppIds)
	if err != nil {
		return nil, err
	}
	var infos []*apps.AppInfo
	if len(appInfos) > 0 {
		for _, info := range appInfos {
			infos = append(infos, &apps.AppInfo{
				Id:        info.Id,
				AppName:   info.AppName,
				AppId:     info.AppId,
				AccountId: info.AccountId,
				PkgName:   info.PkgName,
				Channel:   info.Channel,
				AppType:   info.AppType,
				Tags:      info.Tags,
				CreatedAt: info.CreatedAt.Unix(),
				UpdatedAt: info.UpdatedAt.Unix(),
			})
		}
	}
	return &apps.AppsResp{
		Apps: infos,
	}, nil
}
