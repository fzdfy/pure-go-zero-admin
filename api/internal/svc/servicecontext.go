package svc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/syncx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"pure-go-zero-admin/api/internal/common/errorx"
	"pure-go-zero-admin/api/internal/config"
	"pure-go-zero-admin/api/internal/middleware"

	sys "pure-go-zero-admin/rpc/sys/sysclient"
)

type ServiceContext struct {
	Config   config.Config
	CheckUrl rest.Middleware
	Sys      sys.Sys

	//Redis *redis.Redis
	Cache cache.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	//newRedis := redis.New(c.Redis.Address, redisConfig(c))
	redis := cache.New(c.CacheRedis, syncx.NewSingleFlight(), cache.NewStat("admin-api"), errorx.ErrNotFound)
	return &ServiceContext{
		Config: c,
		//CheckUrl: middleware.NewCheckUrlMiddleware(newRedis).Handle,
		CheckUrl: middleware.NewAuthMiddleware(c, redis).Handle,
		Sys:      sys.NewSys(zrpc.MustNewClient(c.SysRpc)),

		//Redis: newRedis,
		Cache: redis,
	}
}

//func redisConfig(c config.Config) redis.Option {
//	return func(r *redis.Redis) {
//		r.Type = redis.NodeType
//		r.Pass = c.Redis.Pass
//	}
//}
