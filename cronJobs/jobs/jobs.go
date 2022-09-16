package jobs

import (
	"business/common/curl"
	"business/cronJobs/model"
	"business/cronJobs/svc"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

type AccessToken struct {
	Error            int64  `json:"error"`
	ErrorDescription string `json:"error_description"`
	AccessToken      string `json:"access_token"`
	ExpiresIn        int64  `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
}

type QueryParam struct {
	AccountId   int64
	AccessToken string
}

type QueueData interface {
	GenerateMsg(fn func(interface{}))
}

// SetDataToKafka 保存数据到 kafka
func SetDataToKafka(kfk sarama.SyncProducer, q QueueData, kafkaTopic string) error {
	kafkaMsg := make([]*sarama.ProducerMessage, 0)
	q.GenerateMsg(func(data interface{}) {
		bs, err := json.Marshal(data)
		if err != nil {
			return
		}
		kafkaMsg = append(kafkaMsg, &sarama.ProducerMessage{
			Topic: kafkaTopic,
			Value: sarama.ByteEncoder(bs),
		})
	})
	return kfk.SendMessages(kafkaMsg)
}

func GetTokens(ctx context.Context, svcCtx *svc.ServiceContext, tokenChan chan *QueryParam) {
	list, err := svcCtx.TokenModel.GetAccessTokenList(ctx)
	if err != nil {
		return
	}
	for _, tokens := range list {
		if tokens.ExpiredAt.Before(time.Now()) {
			at, err := Refresh(ctx, svcCtx, tokens)
			if err != nil {
				log.Println("Token 刷新失败，账户 ID：", tokens.AccountId, err)
				continue
			}
			tokenChan <- &QueryParam{
				AccountId:   tokens.AccountId,
				AccessToken: fmt.Sprintf("%s %s", tokens.TokenType, at),
			}
		} else {
			tokenChan <- &QueryParam{
				AccountId:   tokens.AccountId,
				AccessToken: fmt.Sprintf("%s %s", tokens.TokenType, tokens.AccessToken),
			}
		}
	}

	close(tokenChan)
}

// Refresh token 过期时刷新
func Refresh(ctx context.Context, svcCtx *svc.ServiceContext, token *model.Tokens) (string, error) {
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
	var at AccessToken
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
		ExpiredAt:    time.Now().Add(time.Duration(at.ExpiresIn - 20)),
		UpdatedAt:    time.Now(),
		TokenType:    at.TokenType,
	})

	return at.AccessToken, nil
}
