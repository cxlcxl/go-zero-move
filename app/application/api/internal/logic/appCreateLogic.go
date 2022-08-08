package logic

import (
	"business/app/application/rpc/appcenter"
	"business/common/utils"
	"business/common/vars"
	"context"

	"business/app/application/api/internal/svc"
	"business/app/application/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppCreateLogic {
	return &AppCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AppCreateLogic) AppCreate(req *types.AppCreateReq) (resp *types.BaseResp, err error) {
	if err = l.svcCtx.Validate.StructCtx(l.ctx, req); err != nil {
		return nil, err
	}
	_, err = l.svcCtx.AppRpcClient.CreateApp(l.ctx, &appcenter.CreateAppReq{
		AppName:   req.AppName,
		AppId:     req.AppId,
		AccountId: req.AccountId,
		PkgName:   req.PkgName,
		Channel:   req.Channel,
		AppType:   req.AppType,
		Tags:      req.Tags,
	})
	if err != nil {
		return nil, utils.RpcError(err, "")
	}
	return &types.BaseResp{
		Code: vars.ResponseCodeOk,
		Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
	}, nil
}
