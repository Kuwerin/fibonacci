package grpc

import (
	"context"

	"github.com/Kuwerin/fibonacci/pkg/service"
	"github.com/Kuwerin/fibonacci/pkg/transport/grpc/fibonaccipb"
)

type GRPCServer struct {
	service *service.Service

	fibonaccipb.UnimplementedFibonacciServiceServer
}

func NewGRPCServer(service *service.Service) *GRPCServer {
	return &GRPCServer{
		service: service,
	}
}

func (s *GRPCServer) GetFibonacciSlice(ctx context.Context, req *fibonaccipb.GetFibonacciSliceRequest) (*fibonaccipb.GetFibonacciSliceResponse, error) {
	return s.service.GetSlice(req)
}
