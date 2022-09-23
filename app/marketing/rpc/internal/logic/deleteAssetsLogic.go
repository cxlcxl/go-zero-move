package logic

import (
	"context"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAssetsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAssetsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAssetsLogic {
	return &DeleteAssetsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteAssetsLogic) DeleteAssets(in *marketing.AssetDeleteReq) (*marketing.BaseResp, error) {
	if err := l.svcCtx.AssetModel.DeleteByAssetIds(l.ctx, in.AssetIds); err != nil {
		return nil, err
	}

	return &marketing.BaseResp{}, nil
}
