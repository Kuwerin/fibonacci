package math_custom

import "github.com/Kuwerin/fibonacci/pkg/cache"

type Fibonacci interface {
	GetSlice(offset, limit uint64) []uint64
}

type MathFeautures struct {
	Fibonacci
}

func NewMathFeautures(cache *cache.Store) *MathFeautures {
	return &MathFeautures{
		Fibonacci: NewFibonacciService(cache),
	}
}
