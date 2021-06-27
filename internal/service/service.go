package service

import (
	"github.com/Kuwerin/fibonacci/internal/model"
	"github.com/Kuwerin/fibonacci/pkg/cache"
)

type Fibonacci interface {
	GetSlice(body model.Fibonacci) []uint64
}


type Service struct {
	Fibonacci
}

func NewService(cache *cache.Store) *Service {
	return &Service{
		Fibonacci: NewFibonacciService(cache),
	}
}