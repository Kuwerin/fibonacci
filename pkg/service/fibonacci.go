package service

import (
	"math"

	"github.com/Kuwerin/fibonacci/pkg/domain"
	"github.com/Kuwerin/fibonacci/pkg/repository"
	"github.com/Kuwerin/fibonacci/pkg/transport/grpc/fibonaccipb"
)

type FibonacciService struct {
	rep *repository.Repository
}

func NewFibonacciService(rep *repository.Repository) *FibonacciService {
	return &FibonacciService{
		rep: rep,
	}
}

func (s FibonacciService) GetSlice(req *fibonaccipb.GetFibonacciSliceRequest) (*fibonaccipb.GetFibonacciSliceResponse, error) {
	var res = new(fibonaccipb.GetFibonacciSliceResponse)

	var fibSlice []uint64
	for i := req.Start; i <= req.End; i++ {
		var fibonacciNum domain.Fibonacci

		fibonacciNum, err := s.rep.Fibonacci.Find(i)
		if err != nil {
			fibonacciNum.Value = countFibonacci(i)

			if err := s.rep.Fibonacci.Save(fibonacciNum); err != nil {
				return nil, err
			}
		}

		fibSlice = append(fibSlice, fibonacciNum.Value)
	}

	res.FibonacciNumber = fibSlice

	return res, nil
}

// count fibonacci via closed formula (https://habr.com/ru/post/261159/)
func countFibonacci(n uint64) uint64 {
	g := (1 + math.Sqrt(5)) / 2

	res := (math.Pow(g, float64(n)) - math.Pow(1-g, float64(n))) / math.Sqrt(5)

	return uint64(res)
}
