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
	"log"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictionaryQueryLogic struct {
	logx.Logger
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	tokenChan chan *QueryParam
	wg        sync.WaitGroup
	statDay   string
	pageSize  int64
}

type QueryParam struct {
	AccountId   int64
	AccessToken string
}

func NewDictionaryQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext, day string) *DictionaryQueryLogic {
	return &DictionaryQueryLogic{
		Logger:    logx.WithContext(ctx),
		ctx:       ctx,
		svcCtx:    svcCtx,
		tokenChan: make(chan *QueryParam),
		wg:        sync.WaitGroup{},
		statDay:   day,
		pageSize:  50,
	}
}

func (l *DictionaryQueryLogic) DictionaryQuery() (err error) {
	go l.getTokens()

	for token := range l.tokenChan {
		l.wg.Add(1)
		go l.query(token, 1)
	}
	l.wg.Wait()
	return
}

func (l *DictionaryQueryLogic) getTokens() {
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

func (l *DictionaryQueryLogic) query(param *QueryParam, page int64) (err error) {
	defer l.wg.Done()
	data := statements.DictionaryRequest{
		Language: "zh_CN",
		TargetingList: []string{
			"pre_define_audience", "not_pre_define_audience", "gender", "age", "series_type", "device_price", "app_category",
			"network_type", "pre_define_audience", "not_pre_define_audience", "media_app_category", "carrier",
			"language", "app_interest", "location_category",
		},
	}
	var response statements.DictionaryResponse
	c, err := curl.New(l.svcCtx.Config.MarketingApis.Tools.Dictionary).Get().JsonData(data)
	if err != nil {
		return err
	}
	if err = c.Request(&response, curl.Authorization(param.AccessToken)); err != nil {
		return err
	}
	if response.Code != "200" {
		return errors.New(response.Message)
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
