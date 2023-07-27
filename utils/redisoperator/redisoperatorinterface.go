package redisoperator

import (
	"context"

	"github.com/go-redis/redis"
)

type IRdb interface {
	Context() context.Context
	Options() *redis.Options
	PSubscribe(channels ...string) *redis.PubSub
	Pipeline() redis.Pipeliner
	Pipelined(fn func(redis.Pipeliner) error) ([]redis.Cmder, error)
	PoolStats() *redis.PoolStats
	SetLimiter(l redis.Limiter) *redis.Client
	Subscribe(channels ...string) *redis.PubSub
	TxPipeline() redis.Pipeliner
	TxPipelined(fn func(redis.Pipeliner) error) ([]redis.Cmder, error)
	Watch(fn func(*redis.Tx) error, keys ...string) error
	WithContext(ctx context.Context) *redis.Client
}

type IOptions interface{}

type IStatusCmd interface {
	Result() (string, error)
	String() string
	Val() string
}
