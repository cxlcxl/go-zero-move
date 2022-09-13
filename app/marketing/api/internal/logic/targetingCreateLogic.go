package logic

import (
	"business/common/vars"
	"context"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TargetingCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTargetingCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TargetingCreateLogic {
	return &TargetingCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TargetingCreateLogic) TargetingCreate(req *types.TargetingCreateReq) (resp *types.TargetingCreateResp, err error) {
	// todo: add your logic here and delete this line

	return &types.TargetingCreateResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: types.TargetingCreateRsInfo{
			TargetingId: "",
		},
	}, nil
}
