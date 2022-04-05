package grpc

import (
	"context"
	"time"

	"github.com/Kuwerin/fibonacci/pkg/service"
	"github.com/Kuwerin/fibonacci/pkg/transport/grpc/fibonaccipb"
	"github.com/go-kit/log"
)

type grpcServer struct {
	service *service.Service
	logger  log.Logger

	fibonaccipb.UnimplementedFibonacciServiceServer
}

func NewGRPCServer(logger log.Logger, service *service.Service) *grpcServer {
	return &grpcServer{
		service: service,
		logger:  logger,
	}
}

func (s *grpcServer) GetFibonacciSlice(ctx context.Context, req *fibonaccipb.GetFibonacciSliceRequest) (res *fibonaccipb.GetFibonacciSliceResponse, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"entity", "transport",
			"type", "grpc",
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.service.GetSlice(req)
}
