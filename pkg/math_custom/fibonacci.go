package math_custom

import (
	"fmt"
	"math"

	"github.com/Kuwerin/fibonacci/pkg/cache"
)

type FibonacciService struct {
	cache *cache.Store
}

func NewFibonacciService(cache *cache.Store) *FibonacciService {
	return &FibonacciService{
		cache: cache,
	}
}

func (s *FibonacciService) GetSlice(offset, limit uint64) []uint64 {
	return s.calcFibonacciSlice(CountFibonacci, offset, limit)
}

type countFibonacciFunc func(uint64) uint64

func (s *FibonacciService) calcFibonacciSlice(countFibonacci countFibonacciFunc, x, y uint64) []uint64 {
	length := y - x + 1
	fibSlice := make([]uint64, length)
	j := 0
	for i := x; i <= y; i++ {
		val, err := s.cache.GetKey(fmt.Sprintf("%d", i))
		if err == nil {
			fibSlice[j] = val
		} else {
			res := countFibonacci(i)
			fibSlice[j] = res
			s.cache.SetKey(fmt.Sprintf("%d", i), res, 0)
		}
		j++
	}
	return fibSlice
}

// count fibonacci via closed formula (https://habr.com/ru/post/261159/)
func CountFibonacci(n uint64) uint64 {
	g := (1 + math.Sqrt(5)) / 2
	res := (math.Pow(g, float64(n)) - math.Pow(1-g, float64(n))) / math.Sqrt(5)
	return uint64(res)
}
