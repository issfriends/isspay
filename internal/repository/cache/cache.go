package cache

import (
	"github.com/issfriends/isspay/internal/app/service"
	"github.com/vx416/gox/cache"
)

var (
	_ service.AuthCacher = (*Cache)(nil)
)

func New(redis *cache.RedisClient) *Cache {
	return &Cache{
		AccountCache: &AccountCache{RedisClient: redis},
	}
}

type Cache struct {
	*AccountCache
}
