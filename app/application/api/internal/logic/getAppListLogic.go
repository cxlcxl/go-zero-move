package logic

import (
	"business/app/application/api/internal/svc"
	"business/app/application/api/internal/types"
	"business/app/application/rpc/appcenter"
	"business/app/marketing/rpc/marketingcenter"
	"business/common/utils"
	"business/common/vars"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAppListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppListLogic {
	return &GetAppListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAppListLogic) GetAppList(req *types.AppListReq) (resp *types.AppListResp, err error) {
	list, err := l.svcCtx.AppRpcClient.AppList(l.ctx, &appcenter.AppListReq{
		AppName:  req.AppName,
		AppId:    req.AppId,
		AppType:  req.AppType,
		Channel:  req.Channel,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, utils.RpcError(err, "数据为空")
	}
	var apps []*types.ListAppInfo
	if list.Total > 0 {
		var actIds []int64
		for _, app := range list.Apps {
			actIds = append(actIds, app.AccountId)
			apps = append(apps, &types.ListAppInfo{
				AppInfo: types.AppInfo{
					Id:        app.Id,
					AppName:   app.AppName,
					AppId:     app.AppId,
					AccountId: app.AccountId,
					PkgName:   app.PkgName,
					Channel:   app.Channel,
					AppType:   app.AppType,
					Tags:      app.Tags,
					CreatedAt: app.CreatedAt,
					UpdatedAt: app.UpdatedAt,
				},
				AccountInfo: types.AccountInfo{},
			})
		}
		accounts, err := l.svcCtx.AccountRpcClient.GetAccountsByAccountIds(l.ctx, &marketingcenter.GetByAccountIdsReq{Ids: actIds})
		if err != nil {
			return nil, utils.RpcError(err, "")
		}
		tmpActMap := make(map[int64]string)
		for _, act := range accounts.Accounts {
			tmpActMap[act.Id] = act.AccountName
		}
		for i := range apps {
			apps[i].AccountName = tmpActMap[apps[i].AccountId]
		}
	}
	return &types.AppListResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		AppChannel: vars.AppChannel,
		Total:      list.Total,
		AppType:    map[int]string{},
		Data:       apps,
	}, nil
}
