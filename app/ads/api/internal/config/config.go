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

	AccountRpcClient zrpc.RpcClientConf

	MarketingApis struct {
		Campaign struct {
			Query  string
			Create string
			Update string
		}

		Adgroup struct {
			Query       string
			Create      string
			Update      string
			BatchStatus string
		}
	}
}
