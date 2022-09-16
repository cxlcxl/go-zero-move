package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Database struct {
		Driver string `yaml:"Driver"`
		Dsn    string `yaml:"Dsn"`
	} `yaml:"Database"`

	MarketingApis struct {
		PageSize int64  `yaml:"PageSize"`
		Refresh  string `yaml:"Refresh"`

		Reports struct {
			CampaignQuery string `yaml:"CampaignQuery"`
			CountryQuery  string `yaml:"CountryQuery"`
		} `yaml:"Reports"`

		Promotion struct {
			Campaign string `yaml:"Campaign"`
		} `yaml:"Promotion"`

		Tools struct {
			Dictionary string `yaml:"Dictionary"`
			Targeting  string `yaml:"Targeting"`
		} `yaml:"Tools"`
	} `yaml:"MarketingApis"`

	Kafka struct {
		Host   string `yaml:"Host"`
		Topics struct {
			Country  string `yaml:"Country"`
			Campaign string `yaml:"Campaign"`
		} `yaml:"Topics"`
	} `yaml:"Kafka"`
}

func Unmarshal(url string, c *Config) error {
	if _, err := os.Stat(url); err != nil {
		log.Fatal("请检查配置文件是否存在：", err)
		return err
	}
	f, err := os.Open(url)
	if err != nil {
		log.Fatal("配置文件打开失败：", err)
		return err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal("配置文件读取失败：", err)
		return err
	}
	err = yaml.UnmarshalStrict(data, c)
	return err
}
