package logic

import (
	"business/common/curl"
	"business/common/kfk"
	"business/cronJobs/jobs"
	"business/cronJobs/model"
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
	tokenChan     chan *QueryParam
	wg            sync.WaitGroup
	statDay       string
	kafkaProducer sarama.SyncProducer
	pageSize      int64
}

type QueryParam struct {
	AccountId   int64
	AccessToken string
}

func NewCampaignQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext, day string) *CampaignQueryLogic {
	return &CampaignQueryLogic{
		Logger:    logx.WithContext(ctx),
		ctx:       ctx,
		svcCtx:    svcCtx,
		tokenChan: make(chan *QueryParam),
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
	go l.getTokens()

	for token := range l.tokenChan {
		l.wg.Add(1)
		go l.query(token, 1)
	}
	l.wg.Wait()
	return
}

func (l *CampaignQueryLogic) getTokens() {
	list, err := l.svcCtx.TokenModel.GetAccessTokenList(l.ctx)
	if err != nil {
		return
	}
	for _, tokens := range list {
		if tokens.ExpiredAt.Before(time.Now()) {
			at, err := refresh(l.ctx, l.svcCtx, tokens)
			if err != nil {
				log.Println("Token 刷新失败，账户 ID：", tokens.AccountId, err)
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
	var at jobs.AccessToken
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

func (l *CampaignQueryLogic) query(param *QueryParam, page int64) (err error) {
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
