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

type TrackingListLogic struct {
	logx.Logger
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	trackings []*types.TrackingItem
}

func NewTrackingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TrackingListLogic {
	return &TrackingListLogic{
		Logger:    logx.WithContext(ctx),
		ctx:       ctx,
		svcCtx:    svcCtx,
		trackings: make([]*types.TrackingItem, 0),
	}
}

func (l *TrackingListLogic) TrackingList(req *types.TrackingListReq) (resp *types.TrackingListResp, err error) {
	trackings, err := l.svcCtx.MarketingRpcClient.TrackingList(l.ctx, &marketingcenter.TrackingListReq{AppId: req.AppId})
	if err != nil && !utils.IsErrNotFound(err) {
		return nil, utils.RpcError(err, "")
	}
	if trackings.Trackings == nil || len(trackings.Trackings) == 0 {
		if err = l.pullTracking(req.AppId); err != nil {
			return nil, err
		}
	} else {
		for _, tracking := range trackings.Trackings {
			l.trackings = append(l.trackings, &types.TrackingItem{
				EffectName: tracking.EffectName,
				EffectType: tracking.EffectType,
				TrackingId: tracking.TrackingId,
			})
		}
	}

	return &types.TrackingListResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: l.trackings,
	}, nil
}

func (l *TrackingListLogic) pullTracking(appId string) (err error) {
	app, err := l.svcCtx.AppRpcClient.GetAppInfoByAppId(l.ctx, &appcenter.GetByAppIdReq{AppId: appId})
	if err != nil {
		return utils.RpcError(err, "应用信息有误")
	}
	token, err := l.svcCtx.AccountRpcClient.GetToken(l.ctx, &accountcenter.GetTokenReq{AccountId: app.AccountId})
	if err != nil {
		return utils.RpcError(err, "Token 信息获取失败")
	}
	if time.Unix(token.ExpiredAt, 0).Before(time.Now()) {
		return errors.New("Token 已过期，请先到账户列表页刷新 Token ")
	}
	t := fmt.Sprintf("%s %s", token.TokenType, token.AccessToken)
	page := 1
	d := statements.TrackingRequest{
		Filtering:    statements.TrackingFilter{ProductUniqueFlag: app.PkgName},
		AdvertiserId: app.AdvertiserId,
		Page:         page,
		PageSize:     50,
	}
	total, err := l.pull(t, d, app.AccountId, app.AppId, app.AdvertiserId)
	if err != nil {
		return err
	}
	if total > 50 {
		pages := ceil(total, 50)
		for i := 2; i <= int(pages); i++ {
			d.Page = i
			if _, err = l.pull(t, d, app.AccountId, app.AppId, app.AdvertiserId); err != nil {
				return err
			}
		}
	}

	return nil
}

func (l *TrackingListLogic) pull(token string, d statements.TrackingRequest, accountId int64, appId, advertiserId string) (int64, error) {
	c, err := curl.New(l.svcCtx.Config.MarketingApis.Tracking.Query).Get().JsonData(d)
	if err != nil {
		return 0, err
	}
	var response statements.TrackingResponse
	if err = c.Request(&response, curl.Authorization(token)); err != nil {
		return 0, err
	}
	if response.Code != "200" {
		return 0, errors.New("华为接口拉取错误：" + response.Message)
	}
	var trackings []*marketingcenter.Tracking
	disabled := make([]int64, 0)
	for _, datum := range response.Data.Data {
		if datum.TrackingStatus == vars.TrackingStatusActive {
			trackings = append(trackings, &marketingcenter.Tracking{
				AppId:            appId,
				AccountId:        accountId,
				AdvertiserId:     advertiserId,
				EffectType:       datum.EffectType,
				EffectName:       datum.EffectName,
				TrackingId:       datum.TrackingId,
				ClickTrackingUrl: datum.ClickTrackingUrl,
				ImpTrackingUrl:   datum.ImpTrackingUrl,
			})
			l.trackings = append(l.trackings, &types.TrackingItem{EffectType: datum.EffectType, EffectName: datum.EffectName, TrackingId: datum.TrackingId})
		} else {
			disabled = append(disabled, datum.TrackingId)
		}
	}
	_, err = l.svcCtx.MarketingRpcClient.BatchInsertTracking(l.ctx, &marketingcenter.BatchInsertTrackingReq{Trackings: trackings, DisabledTrackingIds: disabled})
	if err != nil {
		return 0, utils.RpcError(err, "")
	}
	return response.Data.Total, nil
}
