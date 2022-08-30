package svc

import (
	"business/common/utils"
	"business/cronJobs/jobs/country/config"
	"business/cronJobs/jobs/country/model"
	"github.com/Shopify/sarama"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"log"
)

type ServiceContext struct {
	Config             config.Config
	ReportCountryModel model.ReportCountriesModel
	AdsLogModel        model.AdsRequestLogsModel
	TokenModel         model.TokensModel
	AccountModel       model.AccountsModel
	AppModel           model.AppsModel
	KafkaProducer      sarama.SyncProducer
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := utils.GetDsn(c.Database.User, c.Database.Pass, c.Database.Host, c.Database.DbName, c.Database.Charset, c.Database.Port)
	conn := sqlx.NewSqlConn(c.Database.Driver, dsn)
	return &ServiceContext{
		Config:             c,
		ReportCountryModel: model.NewReportCountriesModel(conn),
		AdsLogModel:        model.NewAdsRequestLogsModel(conn),
		TokenModel:         model.NewTokensModel(conn),
		AccountModel:       model.NewAccountsModel(conn),
		AppModel:           model.NewAppsModel(conn),
		KafkaProducer:      newKafkaProducer(c.Kafka.Host),
	}
}

func newKafkaProducer(kafkaHost string) sarama.SyncProducer {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	kafkaConfig.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	kafkaConfig.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	kafkaProducerClient, err := sarama.NewSyncProducer([]string{kafkaHost}, kafkaConfig)
	if err != nil {
		log.Fatal("new kafka producer error: ", err)
		return nil
	}

	return kafkaProducerClient
}
