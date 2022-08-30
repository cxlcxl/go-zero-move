package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	Auth struct {
		AccessSecret string
		Expire       int64
	}

	Redis struct {
		Host string
		Pass string
	}

	MarketingRpcClient zrpc.RpcClientConf
	AccountRpcClient   zrpc.RpcClientConf

	MarketingApis struct {
		Authorize struct {
			Refresh     string
			CodeUrl     string
			TokenUrl    string
			RedirectUri string
			Scope       string
		}
		Promotion struct {
			Query  string
			Create string
			Update string
		}
	}

	Kafka struct {
		Host   string
		Topics struct {
			Promotion string
		}
	}
}
