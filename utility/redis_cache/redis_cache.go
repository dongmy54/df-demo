package redisCache

import (
	"sync"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gcache"
)

var (
	Cache *gcache.Cache
	once  sync.Once
)

// 封装gcache为redis的缓存
func InitRedis() {
	once.Do(func() {
		redisConfig := &gredis.Config{
			Address: "127.0.0.1:6379",
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
