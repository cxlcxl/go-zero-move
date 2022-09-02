package jobs

import (
	"business/cronJobs/jobs/campaign"
	"business/cronJobs/jobs/country"
	"encoding/json"
	"github.com/Shopify/sarama"
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

type Job func()

var (
	ScheduleJobs = map[string]Job{
		"Country":  country.Country,
		"Campaign": campaign.Campaign,
	}
)

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
