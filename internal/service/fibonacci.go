package service

import (
	"github.com/Kuwerin/fibonacci/internal/model"
	"github.com/Kuwerin/fibonacci/pkg/cache"
	"github.com/Kuwerin/fibonacci/pkg/math_custom"
)

type FibonacciService struct {
	cache *cache.Store
	math  *math_custom.MathFeautures
}

func NewFibonacciService(cache *cache.Store) *FibonacciService {
	return &FibonacciService{
		cache: cache,
		math:  math_custom.NewMathFeautures(cache),
	}
}

func (s FibonacciService) GetSlice(body model.Fibonacci) []uint64 {
	return s.math.GetSlice(body.X, body.Y)
}
