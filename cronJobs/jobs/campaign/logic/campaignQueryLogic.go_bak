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

type CampaignQueryLogic struct {
	logx.Logger
	ctx           context.Context
	svcCtx        *svc.ServiceContext
	tokenChan     chan *jobs.QueryParam
	wg            sync.WaitGroup
	statDay       string
	kafkaProducer sarama.SyncProducer
	pageSize      int64
}

func NewCampaignQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext, day string) *CampaignQueryLogic {
	return &CampaignQueryLogic{
		Logger:    logx.WithContext(ctx),
		ctx:       ctx,
		svcCtx:    svcCtx,
		tokenChan: make(chan *jobs.QueryParam),
		wg:        sync.WaitGroup{},
		statDay:   day,
		pageSize:  50,
	}
}

func (l *CampaignQueryLogic) CampaignQuery() (err error) {
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

func (l *CampaignQueryLogic) query(param *jobs.QueryParam, page int64) (err error) {
	defer l.wg.Done()
	data := statements.CampaignRequest{
		Page:     page,
		PageSize: l.pageSize,
		Filtering: statements.CampaignFiltering{
			UpdatedBeginTime: l.statDay + " 00:00:00",
			UpdatedEndTime:   l.statDay + " 23:59:59",
		},
	}
	var response statements.CampaignResponse
	c, err := curl.New(l.svcCtx.Config.MarketingApis.Promotion.Campaign).Get().JsonData(data)
	if err != nil {
		return err
	}
	if err = c.Request(&response, curl.Authorization(param.AccessToken)); err != nil {
		return err
	}
	if response.Code != "200" {
		return errors.New(response.Message)
	}
	if len(response.Data.Data) == 0 {
		log.Println("查无数据")
		return
	}
	if err = jobs.SetDataToKafka(l.kafkaProducer, response.Data.Data, l.svcCtx.Config.Kafka.Topics.Campaign); err != nil {
		log.Println("数据写入 kafka 失败：", err)
	} else {
		logMsg := fmt.Sprintf("写入 Topic: %s, 数量: %d 条, 页码: %d, 所属账户: %d",
			l.svcCtx.Config.Kafka.Topics.Campaign,
			response.Data.Total,
			page,
			param.AccountId,
		)
		log.Println(logMsg)
	}

	if page > 1 {
		return
	}

	sumPages := ceil(response.Data.Total, l.pageSize)
	if sumPages > 1 {
		var i int64 = 2
		for ; i <= sumPages; i++ {
			if err = l.query(param, i); err != nil {
				fmt.Println("第 ", i, " 页数据拉取出错", err)
			}

			time.Sleep(time.Millisecond * 500)
		}
	}

	return nil
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
