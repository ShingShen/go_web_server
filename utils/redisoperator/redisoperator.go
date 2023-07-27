package redisoperator

import (
	"time"

	"github.com/go-redis/redis"
)

func NewClient(opt *redis.Options) *redis.Client {
	return redis.NewClient(opt)
}

type Redis struct {
	rdb *redis.Client
	// rdb IRdb
}

func (r *Redis) Options() IOptions {
	return r.rdb.Options()
}

func (r *Redis) Set(key string, value interface{}, expiration time.Duration) IStatusCmd {
	return r.rdb.Set(key, value, expiration)
}

func (r *Redis) FlushDB() IStatusCmd {
	return r.rdb.FlushDB()
}

func (r *Redis) Ping() IStatusCmd {
	return r.rdb.Ping()
}
