package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// NewRedis redis模式
func NewRedis(client *redis.Client, options *redis.Options) (*Redis, error) {

	if client == nil {
		client = redis.NewClient(options)
	}
	r := &Redis{
		client: client,
		ctx:    context.Background(),
	}
	err := r.connect()
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Redis cache implement
type Redis struct {
	client *redis.Client
	ctx    context.Context
}

func (*Redis) String() string {
	return "redis"
}

// connect connect test
func (r *Redis) connect() error {
	var err error
	_, err = r.client.Ping(r.ctx).Result()
	return err
}

// Get from key
func (r *Redis) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

// Set value with key and expire time
func (r *Redis) Set(key string, val interface{}, expire int) error {
	return r.client.Set(r.ctx, key, val, time.Duration(expire)*time.Second).Err()
}

// Del delete key in redis
func (r *Redis) Del(key string) error {
	return r.client.Del(r.ctx, key).Err()
}

// HashGet from key
func (r *Redis) HashGet(hk, key string) (string, error) {
	return r.client.HGet(r.ctx, hk, key).Result()
}

// HashDel delete key in specify redis's hashtable
func (r *Redis) HashDel(hk, key string) error {
	return r.client.HDel(r.ctx, hk, key).Err()
}

// Increase
func (r *Redis) Increase(key string) error {
	return r.client.Incr(r.ctx, key).Err()
}

func (r *Redis) Decrease(key string) error {
	return r.client.Decr(r.ctx, key).Err()
}

// Set ttl
func (r *Redis) Expire(key string, dur time.Duration) error {
	return r.client.Expire(r.ctx, key, dur).Err()
}

// GetClient 暴露原生client
func (r *Redis) GetClient() *redis.Client {
	return r.client
}
