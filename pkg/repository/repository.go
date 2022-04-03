package repository

import (
	"github.com/go-redis/redis/v8"

	"github.com/Kuwerin/fibonacci/pkg/repository/fibonacci"
)

type Repository struct {
	Fibonacci fibonacci.Repository
}

func MakeRepository(client *redis.Client) *Repository {
	return &Repository{
		Fibonacci: fibonacci.NewRedisRepository(client),
	}
}
