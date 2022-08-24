package logic

import (
	"business/common/curl"
	"business/common/utils"
	"business/cronJobs/common"
	"business/cronJobs/jobs/country/internal/svc"
	"business/cronJobs/jobs/country/model"
	"business/cronJobs/statements"
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountryQueryLogic struct {
	logx.Logger
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	tokenChan chan *QueryParam
	wg        sync.WaitGroup
	stateDay  string
	appMap    map[string]*app
}

type QueryParam struct {
	AccountId   int64
	AccessToken string
}

type app struct {
	appId   string
	appName string
}

func NewCountryQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountryQueryLogic {
	return &CountryQueryLogic{
		Logger:    logx.WithContext(ctx),
		ctx:       ctx,
		svcCtx:    svcCtx,
		tokenChan: make(chan *QueryParam),
		wg:        sync.WaitGroup{},
		appMap:    map[string]*app{},
	}
}

func (l *CountryQueryLogic) CountryQuery() (err error) {
	if err = l.getApps(); err != nil {
		fmt.Println("======> get app info error: ", err)
		return err
	}

	go l.getTokens()

	err = l.queryCountries()

	l.wg.Wait()
	return
}

func (l *CountryQueryLogic) queryCountries() error {
	for token := range l.tokenChan {
		l.wg.Add(1)
		go l.query(token)
	}

	return nil
}

func (l *CountryQueryLogic) getTokens() {
	list, err := l.svcCtx.TokenModel.GetAccessTokenList(l.ctx)
	if err != nil {
		fmt.Println("======> get token list error: ", err)
		return
	}
	for _, tokens := range list {
		if tokens.ExpiredAt.Before(time.Now()) {
			at, err := refresh(l.ctx, l.svcCtx, tokens)
			if err != nil {
				fmt.Println("======> refresh token error: ", err)
				continue
			}
			l.tokenChan <- &QueryParam{
				AccountId:   tokens.AccountId,
				AccessToken: fmt.Sprintf("%s %s", tokens.TokenType, at),
			}
		} else {
			l.tokenChan <- &QueryParam{
				AccountId:   tokens.AccountId,
				AccessToken: fmt.Sprintf("%s %s", tokens.TokenType, tokens.AccessToken),
			}
		}
	}

	close(l.tokenChan)
}

// token 过期时刷新
func refresh(ctx context.Context, svcCtx *svc.ServiceContext, token *model.Tokens) (string, error) {
	info, err := svcCtx.AccountModel.FindOne(ctx, token.AccountId)
	if err != nil {
		return "", err
	}
	clientId, secret := info.ClientId, info.Secret
	if clientId == "" || secret == "" {
		if info.ParentId == 0 {
			return "", errors.New("未完整填写 ClientId 与 Secret")
		} else {
			parent, err := svcCtx.AccountModel.FindOne(ctx, info.ParentId)
			if err != nil {
				return "", errors.New("上级信息查询错误")
			}
			if parent.ClientId == "" || parent.Secret == "" {
				return "", errors.New("上级 ClientId 与 Secret 信息未填写")
			}
			clientId, secret = parent.ClientId, parent.Secret
		}
	}
	data := map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": token.RefreshToken,
		"client_id":     clientId,
		"client_secret": secret,
	}
	post := curl.New(svcCtx.Config.MarketingApis.Refresh).Post().QueryData(data)
	var at logic.AccessToken
	err = post.Request(&at, curl.FormHeader())
	if err != nil {
		return "", err
	}
	if at.Error != 0 {
		return "", errors.New("华为接口调用失败：" + at.ErrorDescription)
	}
	_ = svcCtx.TokenModel.Update(ctx, &model.Tokens{
		Id:           token.Id,
		AccountId:    token.AccountId,
		AccessToken:  at.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiredAt:    time.Now().Add(time.Duration(at.ExpiresIn) - 20),
		UpdatedAt:    time.Now(),
		TokenType:    at.TokenType,
	})

	return at.AccessToken, nil
}

func (l *CountryQueryLogic) query(param *QueryParam) error {
	defer l.wg.Done()

	var page int64 = 1
	var pageSize int64 = 500
	hasRequestLog := false
	statDate := "2022-08-11" //time.Now().AddDate(0, 0, -2).Format("2006-01-02")
	requestLog, err := l.svcCtx.AdsLogModel.FindLogByStatDayApiModule(l.ctx, statDate, statements.ApiModuleCountry, param.AccountId)
	if err != nil {
		if !errors.Is(err, model.ErrNotFound) {
			fmt.Println("======> get request logs error: ", err)
			return err
		}
	}
	if requestLog != nil {
		if requestLog.IsCompleted == 1 {
			fmt.Println("======> completed <====== ", requestLog)
			return nil
		}
		page = requestLog.NextRequestPage
		hasRequestLog = true
	}
	data := statements.CountryRequest{
		TimeGranularity: statements.StateTimeDaily,
		StartDate:       statDate,
		EndDate:         statDate,
		Page:            page,
		PageSize:        pageSize,
		IsAbroad:        true,
		OrderType:       statements.OrderAsc,
		Filtering: statements.Filtering{
			OtherFilterType: statements.FilterTypeCreative,
		},
	}
	var response statements.CountryResponse
	url := l.svcCtx.Config.MarketingApis.Reports.CountryQuery
	if err = logic.GetHWAdsPostResponse(url, data, &response, param.AccessToken); err != nil {
		return err
	}
	if response.Code != "0" {
		fmt.Println("======> ads api error: ", response.Code, response.Message)
		return errors.New(response.Message)
	}
	var rcs = make([]*model.ReportCountries, len(response.Data.List))
	var t time.Time
	for i, list := range response.Data.List {
		t, err = time.Parse("20060102", list.StatDatetime[:8])
		if err != nil {
			continue
		}
		appId, appName := "", ""
		if tmpApp, ok := l.appMap[list.PackageName]; ok {
			appId, appName = tmpApp.appId, tmpApp.appName
		}
		rcs[i] = &model.ReportCountries{
			StatDate:              t,
			StatHour:              list.StatDatetime[8:],
			AdvertiserId:          list.AdvertiserId,
			AccountId:             param.AccountId,
			CampaignId:            list.CampaignId,
			CampaignName:          strings.TrimSpace(list.CampaignName),
			AdgroupId:             list.AdgroupId,
			AdgroupName:           strings.TrimSpace(list.AdgroupName),
			CreativeId:            list.CreativeId,
			CreativeName:          strings.TrimSpace(list.CreativeName),
			Country:               list.Country,
			AppId:                 appId,
			AppName:               appName,
			ShowCount:             list.ShowCount,
			ClickCount:            list.ClickCount,
			Cpc:                   utils.STF(list.Cpc),
			Cpm:                   utils.STF(list.Cpm),
			Cost:                  utils.STF(list.Cost),
			InstallCount:          list.InstallCount,
			Cpi:                   utils.STF(list.Cpi),
			DownloadCount:         list.DownloadCount,
			Cpd:                   utils.STF(list.Cpd),
			BrowseCount:           list.BrowseCount,
			BrowseCost:            utils.STF(list.BrowseCost),
			AppCustomCount:        list.AppCustomCount,
			AppCustomCost:         utils.STF(list.AppCustomCost),
			WebCustomCount:        list.WebCustomCount,
			WebCustomCost:         utils.STF(list.WebCustomCost),
			PlayCount:             list.PlayCount,
			PlayOverCount:         list.PlayOverCount,
			ThreeDayRetainCount:   list.ThreeDayRetainCount,
			ThreeDayRetainCost:    utils.STF(list.ThreeDayRetainCost),
			ShareCount:            list.ShareCount,
			ShareCost:             utils.STF(list.ShareCost),
			SevenDayRetainCount:   list.SevenDayRetainCount,
			SevenDayRetainCost:    utils.STF(list.SevenDayRetainCost),
			ActiveCountNormalized: list.ActiveCountNormalized,
			Cpa:                   utils.STF(list.Cpa),
			RetainCountNormalized: list.RetainCountNormalized,
			RetainCostNormalized:  utils.STF(list.RetainCostNormalized),
			PayCountNormalized:    list.PayCountNormalized,
			PayCostNormalized:     utils.STF(list.PayCostNormalized),
			CreatedAt:             time.Now(),
			UpdatedAt:             time.Now(),
		}
	}
	adsModel := logic.GetAdsModel(
		statDate,
		statements.ApiModuleCountry,
		param.AccountId,
		page,
		response.Data.PageInfo.TotalPage,
		response.Data.PageInfo.TotalNum,
		pageSize,
		requestLog,
		data,
	)
	if err = l.svcCtx.ReportCountryModel.BatchInsertAndLog(l.ctx, rcs, adsModel, hasRequestLog); err != nil {
		fmt.Println("=======> insert error: ", err)
		return err
	}
	return nil
}

func (l *CountryQueryLogic) getApps() error {
	list, err := l.svcCtx.AppModel.GetApps(l.ctx)
	if err != nil {
		fmt.Println("======> get token list error: ", err)
		return err
	}
	for _, _app := range list {
		l.appMap[_app.PkgName] = &app{
			appId:   _app.AppId,
			appName: _app.AppName,
		}
	}

	return nil
}
