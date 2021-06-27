package server

import (
	"context"
	"net"

	"github.com/Kuwerin/fibonacci/internal/model"
	"github.com/Kuwerin/fibonacci/internal/service"
	pb "github.com/Kuwerin/fibonacci/proto"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// "context"
// "net"

// "github.com/spf13/viper"
// "google.golang.org/grpc"
// "google.golang.org/grpc/reflection"

type GRPCServer struct {
	services *service.Service
	srv *grpc.Server
}

func NewGRPCServer(services *service.Service) *GRPCServer{
	return &GRPCServer{services: services}
}

// func (s *GRPCServer) GetFibSlice(ctx context.Context, req *pb.FibRequest) (*pb.FibResponse, error){
	// res, err := s.services.Fibonacci.GetSlice(model.Fibonacci{
	// 	X: req.X,
	// 	Y: req.Y,
	// })



func (s *GRPCServer) GetFibSlice(ctx context.Context, req *pb.FibRequest) (*pb.FibResponse, error) {
	res := s.services.Fibonacci.GetSlice(model.Fibonacci{
		X: req.A,
		Y: req.B,
	})
	return &pb.FibResponse{Res: res}, nil
}

func (s *GRPCServer) Run(services *service.Service) error {
	l, err := net.Listen("tcp", viper.GetString("grpc.port"))
	if err != nil {
		return err
	}
	s.srv = grpc.NewServer()
	pb.RegisterFibonacciServer(s.srv, &GRPCServer{services: services})
	reflection.Register(s.srv)
	if err := s.srv.Serve(l); err != nil{
		// log.Fatalf("an error occured while trying to listen rpc: %s", err.Error())
		return err
	}
	return nil
}

func (s *GRPCServer) Shutdown()  {
	s.srv.GracefulStop()
}