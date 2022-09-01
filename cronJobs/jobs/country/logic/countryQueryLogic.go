package logic

import (
	"business/common/curl"
	"business/common/kfk"
	"business/cronJobs/jobs/commonLogic"
	"business/cronJobs/model"
	"business/cronJobs/svc"
	"business/cronJobs/vars/statements"
	"context"
	"encoding/json"
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
	tokenChan     chan *QueryParam
	wg            sync.WaitGroup
	statDay       string
	appMap        map[string]*app
	kafkaProducer sarama.SyncProducer
}

type QueryParam struct {
	AccountId   int64
	AccessToken string
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
		tokenChan: make(chan *QueryParam),
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
	go l.getTokens()

	for token := range l.tokenChan {
		l.wg.Add(1)
		go l.query(token, 1)
	}
	l.wg.Wait()
	return
}

func (l *CountryQueryLogic) getTokens() {
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

func (l *CountryQueryLogic) query(param *QueryParam, page int64) (err error) {
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
	if err = logic.GetHWAdsPostResponse(url, data, &response, param.AccessToken); err != nil {
		return err
	}
	if response.Code != "0" {
		return errors.New(response.Message)
	}
	if err = setDataToKafka(l.kafkaProducer, response.Data.List, l.svcCtx.Config.Kafka.Topics.Country); err != nil {
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

// 保存数据到 kafka
func setDataToKafka(kfk sarama.SyncProducer, d []*statements.CountryList, kafkaTopic string) error {
	kafkaMsg := make([]*sarama.ProducerMessage, 0)
	for _, list := range d {
		bs, err := json.Marshal(list)
		if err != nil {
			continue
		}
		kafkaMsg = append(kafkaMsg, &sarama.ProducerMessage{
			Topic: kafkaTopic,
			Value: sarama.ByteEncoder(bs),
		})
	}
	return kfk.SendMessages(kafkaMsg)
}
