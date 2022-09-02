package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Database struct {
		Driver  string
		Host    string
		Port    int64
		User    string
		Pass    string
		DbName  string
		Charset string
	}

	Cache cache.CacheConf

	MarketingApis struct {
		Refresh string
	}
}
