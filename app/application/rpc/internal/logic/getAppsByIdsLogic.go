package logic

import (
	"context"

	"business/app/application/rpc/apps"
	"business/app/application/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppsByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAppsByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppsByIdsLogic {
	return &GetAppsByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAppsByIdsLogic) GetAppsByIds(in *apps.GetByIdsReq) (*apps.AppsResp, error) {
	appInfo, err := l.svcCtx.AppModel.GetAppsByIds(l.ctx, in.Ids)
	if err != nil {
		return nil, err
	}
	var info []*apps.AppInfo
	for _, a := range appInfo {
		info = append(info, &apps.AppInfo{
			Id:        a.Id,
			AppName:   a.AppName,
			AppId:     a.AppId,
			AccountId: a.AccountId,
			PkgName:   a.PkgName,
			Channel:   a.Channel,
			AppType:   a.AppType,
			Tags:      a.Tags,
			CreatedAt: a.CreatedAt.Unix(),
			UpdatedAt: a.UpdatedAt.Unix(),
		})
	}
	return &apps.AppsResp{
		Apps: info,
	}, nil
}
