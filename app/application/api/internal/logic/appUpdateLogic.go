package logic

import (
	"business/app/application/rpc/appcenter"
	"business/common/utils"
	"business/common/vars"
	"context"
	"time"

	"business/app/application/api/internal/svc"
	"business/app/application/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppUpdateLogic {
	return &AppUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AppUpdateLogic) AppUpdate(req *types.AppInfo) (resp *types.BaseResp, err error) {
	if err = l.svcCtx.Validate.StructCtx(l.ctx, req); err != nil {
		return nil, err
	}
	_, err = l.svcCtx.AppRpcClient.UpdateApp(l.ctx, &appcenter.AppInfo{
		Id:        req.Id,
		AppName:   req.AppName,
		AppId:     req.AppId,
		AccountId: req.AccountId,
		PkgName:   req.PkgName,
		Channel:   req.Channel,
		AppType:   req.AppType,
		Tags:      req.Tags,
		UpdatedAt: time.Now().Unix(),
	})
	if err != nil {
		return nil, utils.RpcError(err, "")
	}
	return &types.BaseResp{
		Code: vars.ResponseCodeOk,
		Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
	}, nil
}
