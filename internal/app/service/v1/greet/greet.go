package greet

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	pb "go-scaffold/internal/app/api/scaffold/v1/greet"
	"go-scaffold/internal/app/config"
)

type Service struct {
	pb.UnimplementedGreetServer
	logger *log.Helper
	conf   *config.Config
}

func NewService(logger log.Logger, conf *config.Config) *Service {
	return &Service{
		logger: log.NewHelper(logger),
		conf:   conf,
	}
}

func (s *Service) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Msg: fmt.Sprintf("Hello, %s", req.Name),
	}, nil
}
