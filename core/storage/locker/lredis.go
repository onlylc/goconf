package locker

// import (
// 	"context"
// 	"time"

// 	"github.com/bsm/redislock"
// 	"github.com/redis/go-redis/v9"
// )

// // NewRedis 初始化locker
// func NewRedis(c *redis.Client) *Redis {
// 	return &Redis{
// 		client: c,
// 	}
// }

// var ctx context.Context

// type Redis struct {
// 	client *redis.Client
// 	mutex  *redislock.Client
// }

// func (Redis) String() string {
// 	return "redis"
// }

// func NewOption() {
// 	options * redislock.Options
// 	ctx context.
// }

// func (r *Redis) Lock(key string, ttl int64, options *redislock.Options) (*redislock.Lock, error) {
// 	if r.mutex == nil {
// 		r.mutex = redislock.New(r.client)
// 	}
// 	return r.mutex.Obtain(key, time.Duration(ttl)*time.Second, options)
// }
