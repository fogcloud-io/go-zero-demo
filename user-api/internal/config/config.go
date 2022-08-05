package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpcConf zrpc.RpcClientConf
	DB struct{
		DataSource string
	}
	JwtAuth struct{
		AccessSecret string
		AccessExpire int64
	}
	Cache cache.CacheConf
}
