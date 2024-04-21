package config

import (
	"goconf/core/storage/cache"

	"goconf/core/storage"
)

type Cache struct {
	Redis  *RedisConnectOptions
	Memory interface{}
}

var CacheConfig = new(Cache)

func (e Cache) Setup() (storage.AdapterCache, error) {
	if e.Redis != nil {
		options, err := e.Redis.GetRedisOptions()
		if err != nil {
			return nil, err
		}
		r, err := cache.NewRedis(GetRedisClient(), options)
		if err != nil {
			return nil, err
		}
		if _redis == nil {
			_redis = r.GetClient()
		}
		
		return r, nil

	}
	return cache.NewMemory(), nil
}
