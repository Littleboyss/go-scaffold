package role

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go-scaffold/internal/app/repository/role"
)

type ServiceInterface interface {
	Create(ctx context.Context, req CreateRequest) (*CreateResponse, error)
	Update(ctx context.Context, req UpdateRequest) (*UpdateResponse, error)
	Delete(ctx context.Context, req DeleteRequest) error
	Detail(ctx context.Context, req DetailRequest) (*DetailResponse, error)
	List(ctx context.Context, req ListRequest) (ListResponse, error)
}

var _ ServiceInterface = (*Service)(nil)

type Service struct {
	logger *log.Helper
	repo   role.RepositoryInterface
}

func NewService(
	logger log.Logger,
	repo role.RepositoryInterface,
) *Service {
	return &Service{
		logger: log.NewHelper(logger),
		repo:   repo,
	}
}
