package logic

import (
	"business/common/utils"
	"context"
	"github.com/jinzhu/copier"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssetListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAssetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetListLogic {
	return &AssetListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AssetListLogic) AssetList(in *marketing.AssetListReq) (*marketing.AssetListResp, error) {
	pagination, limit := utils.Pagination(in.Page, in.PageSize)
	list, total, err := l.svcCtx.AssetModel.AssetList(l.ctx, in.AppId, in.AssetType, in.Width, in.Height, pagination, limit)
	if err != nil {
		return nil, err
	}
	var rs []*marketing.Asset
	if err = copier.Copy(&rs, list); err != nil {
		return nil, err
	}
	return &marketing.AssetListResp{
		Assets: rs,
		Total:  total,
	}, nil
}
