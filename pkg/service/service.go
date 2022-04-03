package service

import (
	"github.com/Kuwerin/fibonacci/pkg/repository"
	"github.com/Kuwerin/fibonacci/pkg/transport/grpc/fibonaccipb"
)

type FibonacciServicer interface {
	GetSlice(*fibonaccipb.GetFibonacciSliceRequest) (*fibonaccipb.GetFibonacciSliceResponse, error)
}

type Service struct {
	FibonacciServicer
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		FibonacciServicer: NewFibonacciService(rep),
	}
}
