package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisFeautures interface {
	SetKey(key string, value interface{}, expiresAt time.Duration) error
	GetKey(key string) (uint64, error)
}

type Store struct {
	RedisFeautures
}

func NewRedisStore(rdb *redis.Client, ctx context.Context) *Store {
	return &Store{
		RedisFeautures: NewRedisFeatures(rdb, ctx),
	}
}
