package logic

import (
	"business/app/account/rpc/accountcenter"
	"business/app/marketing/api/internal/statements"
	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"
	"business/app/marketing/rpc/marketingcenter"
	"business/common/curl"
	"business/common/utils"
	"business/common/vars"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

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
	if req.CampaignId == "" {
		return nil, errors.New("参数错误，请先选择计划")
	}
	targeting, err := l.svcCtx.MarketingRpcClient.GetTargetingByName(l.ctx, &marketingcenter.GetTargetingByNameReq{TargetingName: req.TargetingName})
	if err != nil && !utils.IsErrNotFound(err) {
		return nil, utils.RpcError(err, "")
	}
	if targeting != nil {
		return nil, errors.New("定向包名已存在")
	}
	campaign, err := l.svcCtx.MarketingRpcClient.CampaignInfo(l.ctx, &marketingcenter.CampaignInfoReq{CampaignId: req.CampaignId})
	if err != nil {
		return nil, utils.RpcError(err, "")
	}
	data, err := l.formatAdsData(req, campaign.AdvertiserId)
	if err != nil {
		return nil, err
	}
	c, err := curl.New(l.svcCtx.Config.MarketingApis.Targeting.Create).Post().JsonData(data)
	if err != nil {
		return nil, err
	}
	t, err := l.getTokenInfo(campaign.AccountId)
	if err != nil {
		return nil, err
	}
	var rs statements.TargetingCreateResp
	if err = c.Request(&rs, curl.Authorization(t)); err != nil {
		return nil, errors.New("Ads 调用失败：" + err.Error())
	}
	if rs.Code != "200" {
		return nil, errors.New("Ads 调用失败：" + rs.Message)
	}

	_, err = l.svcCtx.MarketingRpcClient.TargetingCreate(l.ctx, l.formatRpcData(req, campaign.AdvertiserId, rs.Data.TargetingId, campaign.AccountId))
	if err != nil {
		return nil, utils.RpcError(err, "")
	}
	return &types.TargetingCreateResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: types.TargetingCreateRsInfo{
			TargetingId: rs.Data.TargetingId,
		},
	}, nil
}

func (l *TargetingCreateLogic) formatAdsData(req *types.TargetingCreateReq, advertiserId string) (rs map[string]interface{}, err error) {
	rs = make(map[string]interface{})
	rs["targeting_name"] = req.TargetingName
	rs["advertiser_id"] = advertiserId
	rs["targeting_type"] = vars.TargetingTypeApp
	// 地域
	if req.Location == "1" {
		if req.LocationType == "current" {
			if len(req.IncludeLocation) > 0 {
				rs["current_custom_location_struct"] = statements.TargetingValue{Value: req.IncludeLocation}
			}
			if len(req.ExcludeLocation) > 0 {
				rs["not_current_custom_location_struct"] = statements.TargetingValue{Value: req.ExcludeLocation}
			}
		} else if req.LocationType == "residence" {
			if len(req.IncludeLocation) > 0 {
				rs["residence_custom_location_struct"] = statements.TargetingValue{Value: req.IncludeLocation}
			}
			if len(req.ExcludeLocation) > 0 {
				rs["not_residence_custom_location_struct"] = statements.TargetingValue{Value: req.ExcludeLocation}
			}
		} else {
			return rs, errors.New("未限制地域类型，或地域选择的类型错误")
		}
	}
	// 性别
	if req.Gender != "" {
		rs["gender_struct"] = statements.TargetingValue{Value: []string{req.Gender}}
	}
	// 年龄
	if len(req.Age) > 0 {
		if _, ok := utils.StringInArray(req.Age, ""); !ok {
			rs["age_struct"] = statements.TargetingValue{Value: req.Age}
		}
	}
	// App 安装
	if req.InstalledApps == "1" {
		rs["installed_apps_struct"] = statements.TargetingValue{Value: []string{"true"}}
	} else {
		rs["not_installed_apps_struct"] = statements.TargetingValue{Value: []string{"true"}}
	}
	// App 行为
	if req.AppCategory != "" {
		switch req.AppCategory {
		case "1":
			rs["app_category_active_struct"] = statements.TargetingValue{Value: req.AppCategories}
			break
		case "2":
			rs["app_category_install_struct"] = statements.TargetingValue{Value: req.AppCategories}
			break
		case "3":
			rs["not_app_category_install_struct"] = statements.TargetingValue{Value: req.AppCategories}
			break
		default:
			return rs, errors.New("App 行为选择错误")
		}
	}
	// App 兴趣
	if req.AppInterest != "" {
		var appInterests []string
		for _, interest := range req.AppInterests {
			appInterests = append(appInterests, interest[len(interest)-1])
		}
		switch req.AppInterest {
		case "1":
			rs["unlimit_app_interest_struct"] = statements.TargetingValue{Value: appInterests}
			break
		case "2":
			rs["normal_app_interest_struct"] = statements.TargetingValue{Value: appInterests}
			break
		case "3":
			rs["high_app_interest_struct"] = statements.TargetingValue{Value: appInterests}
			break
		default:
			return rs, errors.New("App 兴趣选择错误")
		}
	}
	// 设备
	if req.SeriesType == "1" {
		rs["series_type_struct"] = statements.TargetingValue{Value: req.Series}
	}
	// 联网方式
	if len(req.NetworkType) > 0 {
		if _, ok := utils.StringInArray(req.NetworkType, ""); !ok {
			rs["network_type_struct"] = statements.TargetingValue{Value: req.NetworkType}
		}
	}
	// 自定义人群
	if req.Audience == "1" {
		rs["audience_struct"] = statements.TargetingValue{Value: req.Audiences}
		rs["not_audience_struct"] = statements.TargetingValue{Value: req.NotAudience}
	}
	// 媒体类型
	if req.MediaAppCategory == "1" {
		var categories []string
		for _, appCategoryOfMedia := range req.AppCategoryOfMedia {
			categories = append(categories, appCategoryOfMedia[len(appCategoryOfMedia)-1])
		}
		rs["app_category_of_media_struct"] = statements.TargetingValue{Value: categories}
	}
	// 语言
	if req.LanguageCheck == "1" {
		rs["language_struct"] = statements.TargetingValue{Value: req.Language}
	}
	// 运营商
	if req.Carrier == "1" {
		var carriers []string
		for _, carrier := range req.Carriers {
			carriers = append(carriers, carrier[len(carrier)-1])
		}
		rs["carrier_struct"] = statements.TargetingValue{Value: carriers}
	}

	return
}

func (l *TargetingCreateLogic) formatRpcData(req *types.TargetingCreateReq, advertiserId string, targetingId, accountId int64) (targeting *marketingcenter.Targeting) {
	// 地域
	include, exclude, appInterests, mediaCategories, categories, series := "", "", "", "", "", ""
	if req.Location == "1" {
		include = strings.Join(req.IncludeLocation, vars.TargetingDatabaseSeq)
		exclude = strings.Join(req.ExcludeLocation, vars.TargetingDatabaseSeq)
	}
	// App 兴趣
	if req.AppInterest != "" {
		var tmp []string
		for _, interest := range req.AppInterests {
			tmp = append(tmp, interest[len(interest)-1])
		}
		appInterests = strings.Join(tmp, vars.TargetingDatabaseSeq)
	}
	// 媒体类型
	if req.MediaAppCategory == "1" {
		var tmp []string
		for _, appCategoryOfMedia := range req.AppCategoryOfMedia {
			tmp = append(tmp, appCategoryOfMedia[len(appCategoryOfMedia)-1])
		}
		mediaCategories = strings.Join(tmp, vars.TargetingDatabaseSeq)
	}
	// 运营商
	carriers := ""
	if req.Carrier == "1" {
		var tmp []string
		for _, carrier := range req.Carriers {
			tmp = append(tmp, carrier[len(carrier)-1])
		}
		carriers = strings.Join(tmp, vars.TargetingDatabaseSeq)
	}
	// 语言
	language := ""
	if req.LanguageCheck == "1" {
		language = strings.Join(req.Language, vars.TargetingDatabaseSeq)
	}
	// App 行为
	if req.AppCategory != "" {
		categories = strings.Join(req.AppCategories, vars.TargetingDatabaseSeq)
	}
	// 设备
	if req.SeriesType == "1" {
		series = strings.Join(req.Series, vars.TargetingDatabaseSeq)
	}
	// 自定义人群
	audiences, notAudiences := "", ""
	if req.Audience == "1" {
		audiences = strings.Join(req.Audiences, vars.TargetingDatabaseSeq)
		notAudiences = strings.Join(req.NotAudience, vars.TargetingDatabaseSeq)
	}
	targeting = &marketingcenter.Targeting{
		AccountId:          accountId,
		AdvertiserId:       advertiserId,
		TargetingId:        targetingId,
		TargetingName:      req.TargetingName,
		TargetingType:      vars.TargetingTypeApp,
		LocationType:       req.LocationType,
		IncludeLocation:    include,
		ExcludeLocation:    exclude,
		Carriers:           carriers,
		Language:           language,
		Age:                strings.Join(req.Age, vars.TargetingDatabaseSeq),
		Gender:             req.Gender,
		AppCategory:        req.AppCategory,
		AppCategories:      categories,
		InstalledApps:      req.InstalledApps,
		AppInterest:        req.AppInterest,
		AppInterests:       appInterests,
		Series:             series,
		NetworkType:        strings.Join(req.NetworkType, vars.TargetingDatabaseSeq),
		NotAudiences:       notAudiences,
		Audiences:          audiences,
		AppCategoryOfMedia: mediaCategories,
		CreatedAt:          time.Now().Unix(),
		UpdatedAt:          time.Now().Unix(),
	}

	return
}

func (l *TargetingCreateLogic) getTokenInfo(accountId int64) (token string, err error) {
	tokenInfo, err := l.svcCtx.AccountRpcClient.GetToken(l.ctx, &accountcenter.GetTokenReq{AccountId: accountId})
	if err != nil {
		return "", utils.RpcError(err, "TOKEN 数据异常")
	}
	token = fmt.Sprintf("%s %s", tokenInfo.TokenType, tokenInfo.AccessToken)
	if time.Unix(tokenInfo.ExpiredAt, 0).Before(time.Now()) {
		if tokenInfo, err = l.svcCtx.AccountRpcClient.RefreshToken(l.ctx, &accountcenter.GetTokenReq{AccountId: accountId}); err != nil {
			return "", utils.RpcError(err, "")
		}
	}

	return fmt.Sprintf("%s %s", tokenInfo.TokenType, tokenInfo.AccessToken), nil
}
