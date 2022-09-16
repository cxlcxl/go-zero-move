package logic

import (
	"business/common/curl"
	"business/cronJobs/jobs"
	"business/cronJobs/model"
	"business/cronJobs/svc"
	"business/cronJobs/vars/statements"
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
	"time"
)

type TargetingLogic struct {
	logx.Logger
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	tokenChan chan *jobs.QueryParam
	wg        sync.WaitGroup
	pageSize  int64
}

func NewTargetingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TargetingLogic {
	return &TargetingLogic{
		Logger:    logx.WithContext(ctx),
		ctx:       ctx,
		svcCtx:    svcCtx,
		tokenChan: make(chan *jobs.QueryParam),
		wg:        sync.WaitGroup{},
		pageSize:  50,
	}
}

func (l *TargetingLogic) TargetingQuery() (err error) {
	go jobs.GetTokens(l.ctx, l.svcCtx, l.tokenChan)

	for token := range l.tokenChan {
		account, err := l.svcCtx.AccountModel.FindOne(l.ctx, token.AccountId) // 只有几条数据
		if err != nil {
			fmt.Println("定向包调度失败：", err)
		}
		if err = l.query(token, account.AdvertiserId, 1); err != nil {
			fmt.Println("定向包调度失败：", err)
		}
	}

	return
}

func (l *TargetingLogic) query(param *jobs.QueryParam, advertiserId string, page int64) (err error) {
	data := statements.TargetingRequest{
		RequestPage:  statements.RequestPage{Page: page, PageSize: l.pageSize},
		AdvertiserId: advertiserId,
	}
	c, err := curl.New(l.svcCtx.Config.MarketingApis.Tools.Dictionary).Get().JsonData(data)
	if err != nil {
		return err
	}
	var response statements.TargetingResponse
	if err = c.Request(&response, curl.Authorization(param.AccessToken)); err != nil {
		return err
	}
	if response.Code != "200" {
		return errors.New("接口返回错误：" + response.Message)
	}
	targetings := formatTargeting(response.Data.Targetings, param.AccountId, advertiserId)
	if err = l.svcCtx.TargetingModel.BatchInsert(l.ctx, targetings); err != nil {
		fmt.Println("写入数据错误：", err)
	}
	offset := page * l.pageSize
	if response.Data.Total > offset {
		page++
		if err = l.query(param, advertiserId, page); err != nil {
			fmt.Println("定向包分页调度失败：", page, err)
		}
	}
	return err
}

func formatTargeting(src []*statements.Targeting, accountId int64, advertiserId string) (targetings []*model.Targetings) {
	for _, targeting := range src {
		targetings = append(targetings, &model.Targetings{
			AccountId:          accountId,
			AdvertiserId:       advertiserId,
			TargetingId:        targeting.TargetingId,
			TargetingName:      targeting.TargetingName,
			TargetingType:      targeting.TargetingType,
			LocationType:       "",
			IncludeLocation:    "",
			ExcludeLocation:    "",
			Carriers:           "",
			Language:           "",
			Age:                "",
			Gender:             "",
			AppCategory:        "",
			AppCategories:      "",
			InstalledApps:      "",
			AppInterest:        "",
			AppInterests:       "",
			Series:             "",
			NetworkType:        "",
			NotAudiences:       "",
			Audiences:          "",
			AppCategoryOfMedia: "",
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		})
	}
	return
}
