package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Database struct {
		Driver  string `yaml:"Driver"`
		Host    string `yaml:"Host"`
		Port    int64  `yaml:"Port"`
		User    string `yaml:"User"`
		Pass    string `yaml:"Pass"`
		DbName  string `yaml:"DbName"`
		Charset string `yaml:"Charset"`
	} `yaml:"Database"`

	MarketingApis struct {
		Refresh string `yaml:"Refresh"`

		Reports struct {
			CampaignQuery string `yaml:"CampaignQuery"`
			CountryQuery  string `yaml:"CountryQuery"`
		} `yaml:"Reports"`
	} `yaml:"MarketingApis"`
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
