package logic

import (
	"business/app/account/rpc/accountcenter"
	"business/app/application/rpc/appcenter"
	"business/app/marketing/api/internal/statements"
	"business/app/marketing/rpc/marketingcenter"
	"business/common/curl"
	"business/common/utils"
	"business/common/vars"
	"context"
	"errors"
	"fmt"
	"time"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncAssetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSyncAssetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncAssetLogic {
	return &SyncAssetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncAssetLogic) SyncAsset(req *types.SyncAssetReq) (resp *types.BaseResp, err error) {
	app, err := l.svcCtx.AppRpcClient.GetAppInfoByAppId(l.ctx, &appcenter.GetByAppIdReq{AppId: req.AppId})
	if err != nil {
		return nil, utils.RpcError(err, "应用数据异常")
	}
	token, err := l.svcCtx.AccountRpcClient.GetToken(l.ctx, &accountcenter.GetTokenReq{AccountId: app.AccountId})
	if err != nil {
		return nil, utils.RpcError(err, "Token 信息获取失败")
	}
	if time.Unix(token.ExpiredAt, 0).Before(time.Now()) {
		return nil, errors.New("Token 已过期，请先到账户列表页刷新 Token ")
	}
	page := 1
	d := statements.AdsAssetRequest{
		AdvertiserId: app.AdvertiserId,
		Page:         page,
		PageSize:     50,
		Filtering:    statements.Filtering{AssetStatus: "CREATIVE_ASSET_ENABLE"},
	}
	t := fmt.Sprintf("%s %s", token.TokenType, token.AccessToken)
	total, err := l.pull(t, d, app.AccountId, app.AdvertiserId)
	if err != nil {
		return nil, err
	}

	if total > 50 {
		pages := ceil(total, 50)
		for i := 2; i <= int(pages); i++ {
			d.Page = i
			if _, err = l.pull(t, d, app.AccountId, app.AdvertiserId); err != nil {
				return nil, err
			}
		}
	}
	return &types.BaseResp{
		Code: vars.ResponseCodeOk,
		Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
	}, nil
}

func (l *SyncAssetLogic) pull(token string, d statements.AdsAssetRequest, accountId int64, advertiserId string) (int64, error) {
	c, err := curl.New(l.svcCtx.Config.MarketingApis.Asset.Query).Get().JsonData(d)
	if err != nil {
		return 0, err
	}
	var response statements.AdsAssetResponse
	if err = c.Request(&response, curl.Authorization(token)); err != nil {
		return 0, err
	}
	if response.Code != "200" {
		return 0, errors.New("华为接口拉取错误：" + response.Message)
	}
	assets := make([]*marketingcenter.Asset, len(response.Data.CreativeAssetInfos))
	for i, datum := range response.Data.CreativeAssetInfos {
		assets[i] = &marketingcenter.Asset{
			AccountId:         accountId,
			AdvertiserId:      advertiserId,
			AssetId:           datum.AssetId,
			AssetName:         datum.AssetName,
			AssetType:         datum.AssetType,
			FileUrl:           datum.FileUrl,
			Width:             datum.Width,
			Height:            datum.Height,
			VideoPlayDuration: datum.VideoPlayDuration,
			FileSize:          datum.FileSize,
			FileFormat:        datum.FileFormat,
			FileHashSha256:    datum.FileHashSha256,
		}
	}
	_, err = l.svcCtx.MarketingRpcClient.BatchInsertAsset(l.ctx, &marketingcenter.BatchInsertAssetReq{Assets: assets})
	if err != nil {
		return 0, utils.RpcError(err, "")
	}
	return response.Data.Total, nil
}

func ceil(num, pageSize int64) int64 {
	if num < pageSize {
		return 0
	}
	var d int64 = 0
	if num%pageSize > 0 {
		d = 1
	}
	return num/pageSize + d
}
