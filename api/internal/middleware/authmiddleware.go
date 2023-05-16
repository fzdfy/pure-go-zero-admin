package middleware

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"pure-go-zero-admin/api/internal/common/errorx"
	"pure-go-zero-admin/api/internal/common/prefix"
	"pure-go-zero-admin/api/internal/config"
	"pure-go-zero-admin/util/gconv"
	"pure-go-zero-admin/util/jwtx"
	"time"
)

type AuthMiddleware struct {
	Config config.Config
	Cache  cache.Cache
}

func NewAuthMiddleware(Config config.Config, Cache cache.Cache) *AuthMiddleware {
	//redis := cache.New(c.CacheRedis, syncx.NewSingleFlight(), cache.NewStat("admin-api"), errorx.ErrNotFound)
	return &AuthMiddleware{Config: Config, Cache: Cache}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("userId")
		if userId == "" {
			logx.Errorf("缺少必要参数userId")
			httpx.Error(w, errorx.NewDefaultError("缺少必要参数userId"))
			return
		}

		// 从redis获取token
		headerToken := r.Header.Get("Authorization")
		token := ""
		cacheKey := fmt.Sprintf("%s%v", prefix.CacheSysUserTokenPrefix, userId)
		err := m.Cache.Get(cacheKey, &token)
		if err != nil {
			httpx.Error(w, errorx.NewCodeError(401, "未登录"))
			return
		}
		// 解析token
		claims, _ := jwtx.ParseToken(headerToken, m.Config.Auth.AccessSecret)
		expTime := gconv.Int64(claims["exp"])
		// 判断是否已过期
		if expTime-time.Now().Unix() <= 0 {
			httpx.Error(w, errorx.NewCodeError(401, "登录已过期"))
			return
		}
		next(w, r)
	}
}
