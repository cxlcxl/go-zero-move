package logic

import (
	"business/app/application/rpc/appcenter"
	"business/common/utils"
	"business/common/vars"
	"context"
	"github.com/jinzhu/copier"

	"business/app/application/api/internal/svc"
	"business/app/application/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAppInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppInfoLogic {
	return &GetAppInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAppInfoLogic) GetAppInfo(req *types.AppInfoReq) (resp *types.AppInfoResp, err error) {
	info, err := l.svcCtx.AppRpcClient.GetAppInfo(l.ctx, &appcenter.AppInfoReq{Id: req.Id})
	if err != nil {
		return nil, utils.RpcError(err, "未查询到应用")
	}
	var app types.AppInfo
	if err = copier.Copy(&app, info); err != nil {
		return nil, vars.DataCopierError
	}
	return &types.AppInfoResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: app,
	}, nil
}
