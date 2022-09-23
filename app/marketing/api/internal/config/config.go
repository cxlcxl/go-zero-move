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
	AppRpcClient       zrpc.RpcClientConf

	MarketingApis struct {
		Authorize struct {
			Refresh string
		}
		Campaign struct {
			Query  string
			Create string
			Update string
		}
		Targeting struct {
			Query  string
			Create string
			Update string
		}
		Asset struct {
			Query        string
			Delete       string
			Create       string
			Token        string
			AssetTmpPath string
		}
	}

	Kafka struct {
		Host   string
		Topics struct {
			Promotion string
		}
	}
}
