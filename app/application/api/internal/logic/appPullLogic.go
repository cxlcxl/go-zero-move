package logic

import (
	"business/app/account/rpc/accountcenter"
	"business/app/application/api/internal/statements"
	"business/app/application/api/internal/svc"
	"business/app/application/api/internal/types"
	"business/app/application/rpc/appcenter"
	"business/common/curl"
	"business/common/utils"
	"business/common/vars"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppPullLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppPullLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppPullLogic {
	return &AppPullLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AppPullLogic) AppPull(req *types.AppPullReq) (resp *types.BaseResp, err error) {
	token, err := l.svcCtx.AccountRpcClient.GetToken(l.ctx, &accountcenter.GetTokenReq{AccountId: req.AccountId})
	if err != nil {
		return nil, errors.New("账户 Token 获取失败，请确认账户是否已认证")
	}
	if time.Unix(token.ExpiredAt, 0).Before(time.Now()) {
		return nil, errors.New("Token 已过期，请先到账户列表页刷新 Token ")
	}
	page := 1
	d := statements.AppPullRequest{
		AdvertiserId: req.AdvertiserId,
		Page:         page,
		PageSize:     50,
		Filtering: statements.ReqFiltering{
			ProductType: vars.ProductTypeAndroidApp,
			AgAppType:   "AG_APP_FOR_DISPLAY_NETWORK",
		},
	}
	t := fmt.Sprintf("%s %s", token.TokenType, token.AccessToken)
	total, err := l.pull(t, d, req.AccountId)
	if err != nil {
		return nil, err
	}

	if total > 50 {
		pages := ceil(total, 50)
		for i := 2; i <= int(pages); i++ {
			d.Page = i
			if _, err = l.pull(t, d, req.AccountId); err != nil {
				return nil, err
			}
		}
	}
	return &types.BaseResp{
		Code: vars.ResponseCodeOk,
		Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
	}, nil
}

func (l *AppPullLogic) pull(token string, d statements.AppPullRequest, accountId int64) (int64, error) {
	c, err := curl.New(l.svcCtx.Config.MarketingApis.App.Query).Get().JsonData(d)
	if err != nil {
		return 0, err
	}
	var response statements.AppAdsResponse
	if err = c.Request(&response, curl.Authorization(token)); err != nil {
		return 0, err
	}
	if response.Code != "200" {
		return 0, errors.New("华为接口拉取错误：" + response.Message)
	}
	apps := make([]*appcenter.Apps, len(response.Data.Data))
	for i, datum := range response.Data.Data {
		apps[i] = &appcenter.Apps{
			AdvertiserId:        d.AdvertiserId,
			AppName:             datum.ProductInfo.App.ProductName,
			AppId:               datum.ProductInfo.App.AppId,
			AccountId:           accountId,
			PkgName:             datum.ProductInfo.App.PackageName,
			AppType:             datum.ProductType,
			IconUrl:             datum.ProductInfo.App.IconUrl,
			ProductId:           datum.ProductId,
			AppStoreDownloadUrl: datum.ProductInfo.App.AppStoreDownloadUrl,
		}
	}
	_, err = l.svcCtx.AppRpcClient.BatchCreateApp(l.ctx, &appcenter.BatchCreateAppReq{Apps: apps})
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
