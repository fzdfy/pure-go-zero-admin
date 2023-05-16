package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	CacheRedis cache.ClusterConf
	//系统
	SysRpc zrpc.RpcClientConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	//Redis struct {
	//	Address string
	//	Pass    string
	//}
}
