package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Features struct {
	ctx context.Context
	rdb redis.Client	
}


func NewRedisFeatures(rdb *redis.Client, ctx context.Context) *Features{
	return &Features{
		ctx: ctx,
		rdb: *rdb,
	}
}

func (r *Features) SetKey(key string, value interface{}, expiresAt time.Duration) error {
	return r.rdb.Set(r.ctx, key, value, expiresAt).Err()
}

func (r *Features) GetKey(key string) (uint64, error) {
	return r.rdb.Get(r.ctx, key).Uint64()
}