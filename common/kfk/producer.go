package kfk

import (
	"github.com/Shopify/sarama"
)

// NewProducer 创建生产者
func NewProducer(kafkaHost string) (sarama.SyncProducer, error) {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	kafkaConfig.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	kafkaConfig.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	kafkaProducerClient, err := sarama.NewSyncProducer([]string{kafkaHost}, kafkaConfig)
	if err != nil {
		return nil, err
	}

	return kafkaProducerClient, nil
}
