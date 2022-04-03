package repository

import (
	"github.com/go-kit/kit/log"
	"github.com/go-redis/redis/v8"

	"github.com/Kuwerin/fibonacci/pkg/repository/fibonacci"
)

type Repository struct {
	Fibonacci fibonacci.Repository
}

func MakeRepository(logger log.Logger, client *redis.Client) *Repository {
	var r = new(Repository)

	r.Fibonacci = fibonacci.NewRedisRepository(client)
	r.Fibonacci = fibonacci.LoggingMiddleware(logger)(r.Fibonacci)

	return r
}
