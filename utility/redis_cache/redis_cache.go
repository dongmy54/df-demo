package redisCache

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

var (
	Cache *gcache.Cache
	once  sync.Once
)

// 封装gcache为redis的缓存
func InitRedis() {
	once.Do(func() {
		ctx := context.TODO()
		addr := g.Cfg().MustGet(ctx, "redis.default.address").String()
		redisConfig := &gredis.Config{
			Address: addr,
			Db:      9,
		}

		// Create redis client object.
		redis, err := gredis.New(redisConfig)
		if err != nil {
			panic(err)
		}
		Cache = gcache.New()
		Cache.SetAdapter(gcache.NewAdapterRedis(redis))
	})
}
