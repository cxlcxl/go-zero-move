package logic

import (
	"business/common/curl"
	"business/common/kfk"
	"business/cronJobs/jobs"
	"business/cronJobs/svc"
	"business/cronJobs/vars/statements"
	"context"
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountryQueryLogic struct {
	logx.Logger
	ctx           context.Context
	svcCtx        *svc.ServiceContext
	tokenChan     chan *jobs.QueryParam
	wg            sync.WaitGroup
	statDay       string
	appMap        map[string]*app
	kafkaProducer sarama.SyncProducer
}

type app struct {
	appId   string
	appName string
}

func NewCountryQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext, day string) *CountryQueryLogic {
	return &CountryQueryLogic{
		Logger:    logx.WithContext(ctx),
		ctx:       ctx,
		svcCtx:    svcCtx,
		tokenChan: make(chan *jobs.QueryParam),
		wg:        sync.WaitGroup{},
		appMap:    map[string]*app{},
		statDay:   day,
	}
}

func (l *CountryQueryLogic) CountryQuery() (err error) {
	if err = l.getApps(); err != nil {
		return err
	}
	l.kafkaProducer, err = kfk.NewProducer(l.svcCtx.Config.Kafka.Host)
	if err != nil {
		return err
	}
	defer l.kafkaProducer.Close()
	go jobs.GetTokens(l.ctx, l.svcCtx, l.tokenChan)

	for token := range l.tokenChan {
		l.wg.Add(1)
		go l.query(token, 1)
	}
	l.wg.Wait()
	return
}

func (l *CountryQueryLogic) query(param *jobs.QueryParam, page int64) (err error) {
	defer l.wg.Done()
	data := statements.CountryRequest{
		TimeGranularity: statements.StateTimeDaily,
		StartDate:       l.statDay,
		EndDate:         l.statDay,
		Page:            page,
		PageSize:        l.svcCtx.Config.MarketingApis.PageSize,
		IsAbroad:        true,
		OrderType:       statements.OrderAsc,
		Filtering: statements.Filtering{
			OtherFilterType: statements.FilterTypeCreative,
		},
	}
	var response statements.CountryResponse
	url := l.svcCtx.Config.MarketingApis.Reports.CountryQuery
	c, err := curl.New(url).Post().JsonData(data)
	if err != nil {
		return err
	}
	if err = c.Request(&response); err != nil {
		return err
	}
	if response.Code != "0" {
		return errors.New(response.Message)
	}
	if err = jobs.SetDataToKafka(l.kafkaProducer, response.Data.List, l.svcCtx.Config.Kafka.Topics.Country); err != nil {
		log.Println("数据写入 kafka 失败：", err)
	} else {
		logMsg := fmt.Sprintf("写入 Topic: %s, 数量: %d 条, 页码: %d, 所属账户: %d",
			l.svcCtx.Config.Kafka.Topics.Country,
			len(response.Data.List),
			page,
			param.AccountId,
		)
		log.Println(logMsg)
	}

	if page > 1 {
		return
	}
	if response.Data.PageInfo.TotalPage > 1 {
		var i int64 = 2
		for ; i <= response.Data.PageInfo.TotalPage; i++ {
			if err = l.query(param, i); err != nil {
				fmt.Println("第 ", i, " 页数据拉取出错", err)
			}

			time.Sleep(time.Millisecond * 500)
		}
	}

	return nil
}

func (l *CountryQueryLogic) getApps() error {
	list, err := l.svcCtx.AppModel.GetApps(l.ctx)
	if err != nil {
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
