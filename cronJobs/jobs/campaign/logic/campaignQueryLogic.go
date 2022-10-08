package logic

import (
	"business/common/curl"
	"business/cronJobs/jobs"
	"business/cronJobs/model"
	"business/cronJobs/svc"
	"business/cronJobs/vars/statements"
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type CampaignQueryLogic struct {
	logx.Logger
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	tokenChan chan *jobs.QueryParam
	statDay   string
	pageSize  int64
	campaigns []*model.Campaigns
	worker    int
}

func NewCampaignQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext, day string) *CampaignQueryLogic {
	return &CampaignQueryLogic{
		Logger:    logx.WithContext(ctx),
		ctx:       ctx,
		svcCtx:    svcCtx,
		tokenChan: make(chan *jobs.QueryParam),
		campaigns: make([]*model.Campaigns, 0),
		statDay:   day,
		pageSize:  50,
		worker:    0,
	}
}

func (l *CampaignQueryLogic) CampaignQuery() (err error) {
	if err != nil {
		return err
	}
	go jobs.GetTokens(l.ctx, l.svcCtx, l.tokenChan)

	for token := range l.tokenChan {
		l.query(token, 1)
	}

	if len(l.campaigns) > 0 {
		err = l.svcCtx.CampaignModel.BatchInsert(l.ctx, l.campaigns)
	}
	return
}

func (l *CampaignQueryLogic) query(param *jobs.QueryParam, page int64) {
	data := statements.CampaignRequest{
		AdvertiserId: param.AdvertiserId,
		Page:         page,
		PageSize:     l.pageSize,
		Filtering: statements.CampaignFiltering{
			UpdatedBeginTime: l.statDay + " 00:00:00",
			UpdatedEndTime:   l.statDay + " 23:59:59",
		},
	}
	var response statements.CampaignResponse
	c, err := curl.New(l.svcCtx.Config.MarketingApis.Promotion.Campaign).Get().JsonData(data)
	if err != nil {
		fmt.Println("参数生成失败", err)
		return
	}
	if err = c.Request(&response, curl.Authorization(param.AccessToken)); err != nil {
		fmt.Println("接口请求失败", err)
		return
	}
	if response.Code != "200" {
		fmt.Println("接口请求失败", response.Code, response.Message)
		return
	}
	if len(response.Data.Data) == 0 {
		return
	}
	now := time.Now()
	for _, datum := range response.Data.Data {
		l.campaigns = append(l.campaigns, &model.Campaigns{
			CampaignId:                datum.CampaignId,
			AppId:                     "",
			CampaignName:              datum.CampaignName,
			AccountId:                 param.AccountId,
			AdvertiserId:              param.AdvertiserId,
			OptStatus:                 datum.CampaignStatus,
			CampaignDailyBudgetStatus: datum.CampaignDailyBudgetStatus,
			ProductType:               datum.ProductType,
			ShowStatus:                datum.ShowStatus,
			UserBalanceStatus:         datum.UserBalanceStatus,
			FlowResource:              datum.FlowResource,
			SyncFlowResource:          datum.SyncFlowResourceSearchAd,
			CampaignType:              datum.CampaignType,
			TodayDailyBudget:          datum.TodayDailyBudget,
			TomorrowDailyBudget:       datum.TomorrowDailyBudget,
			MarketingGoal:             datum.MarketingGoal,
			CreatedAt:                 now,
			UpdatedAt:                 now,
		})
	}

	if page > 1 {
		return
	}

	sumPages := ceil(response.Data.Total, l.pageSize)
	if sumPages > 1 {
		var i int64 = 2
		for ; i <= sumPages; i++ {
			l.query(param, i)

			time.Sleep(time.Millisecond * 500)
		}
	}

	return
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
