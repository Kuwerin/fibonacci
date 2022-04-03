package fibonacci

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/Kuwerin/fibonacci/pkg/domain"
)

type redisRepository struct {
	rdb redis.Client
}

func NewRedisRepository(rdb *redis.Client) *redisRepository {
	return &redisRepository{
		rdb: *rdb,
	}
}

func (r *redisRepository) Save(fibonacciNumber domain.Fibonacci) error {
	return r.rdb.Set(context.Background(), fmt.Sprintf("%d", fibonacciNumber.Key), fibonacciNumber.Value, 0*time.Second).Err()
}

func (r *redisRepository) Find(key uint64) (domain.Fibonacci, error) {
	var fibonacciNumber domain.Fibonacci

	value, err := r.rdb.Get(context.Background(), fmt.Sprintf("%d", key)).Uint64()

	if err != nil {
		return fibonacciNumber, err
	}

	fibonacciNumber.Key = key
	fibonacciNumber.Value = value

	return fibonacciNumber, nil
}
