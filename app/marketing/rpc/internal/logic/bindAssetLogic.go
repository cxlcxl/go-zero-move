package logic

import (
	"context"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindAssetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBindAssetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindAssetLogic {
	return &BindAssetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BindAssetLogic) BindAsset(in *marketing.AssetBindReq) (*marketing.BaseResp, error) {
	if err := l.svcCtx.AssetModel.AssetBind(l.ctx, in.AppId, in.AssetId); err != nil {
		return nil, err
	}
	return &marketing.BaseResp{}, nil
}
