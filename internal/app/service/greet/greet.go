package greet

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/log"
)

type ServiceInterface interface {
	Hello(ctx context.Context, request HelloRequest) (*HelloResponse, error)
}

var _ ServiceInterface = (*Service)(nil)

type Service struct {
	enforcer *casbin.Enforcer
	logger *log.Helper
}

func NewService(enforcer *casbin.Enforcer, logger log.Logger) *Service {
	return &Service{
		enforcer :enforcer,
		logger: log.NewHelper(logger),
	}
}
