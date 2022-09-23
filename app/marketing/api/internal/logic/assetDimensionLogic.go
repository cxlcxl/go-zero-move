package logic

import (
	"business/app/marketing/api/internal/statements"
	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"
	"business/common/vars"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssetDimensionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssetDimensionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetDimensionLogic {
	return &AssetDimensionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssetDimensionLogic) AssetDimension() (resp *types.AssetDimensionResp, err error) {
	var (
		pictureDimension []string
		videoDimension   []string
	)
	for _, dimension := range statements.PictureDimensions {
		pictureDimension = append(pictureDimension, fmt.Sprintf("%d*%d", dimension.Width, dimension.Height))
	}
	for _, dimension := range statements.VideoDimensions {
		videoDimension = append(videoDimension, fmt.Sprintf("%d*%d", dimension.Width, dimension.Height))
	}
	return &types.AssetDimensionResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: map[string][]string{
			vars.AssetTypePicture: pictureDimension,
			vars.AssetTypeVideo:   videoDimension,
		},
	}, nil
}
