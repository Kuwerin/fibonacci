package fibonacci

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"github.com/Kuwerin/fibonacci/pkg/domain"
)

type redisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) *redisRepository {
	return &redisRepository{
		client: client,
	}
}

func (r *redisRepository) Save(fibonacciNumber domain.Fibonacci) error {
	return r.client.Set(context.Background(), fmt.Sprintf("%d", fibonacciNumber.Key), fibonacciNumber.Value, 0).Err()
}

func (r *redisRepository) Find(key uint64) (domain.Fibonacci, error) {
	var fibonacciNumber domain.Fibonacci

	value, err := r.client.Get(context.Background(), fmt.Sprintf("%d", key)).Uint64()

	if err != nil {
		return fibonacciNumber, err
	}

	fibonacciNumber.Key = key
	fibonacciNumber.Value = value

	return fibonacciNumber, nil
}
