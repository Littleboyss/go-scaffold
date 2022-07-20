package greet

import (
	"context"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	errorsx "go-scaffold/internal/app/pkg/errors"
)

type HelloRequest struct {
	Name string `json:"name" form:"name"`
}

func (r HelloRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required.Error("名称不能为空")),
	)
}

type HelloResponse struct {
	Msg string `json:"msg"`
}

func (s *Service) Hello(ctx context.Context, req HelloRequest) (*HelloResponse, error) {
	added, _ := s.enforcer.AddPolicy("Tom", "/v1/greet", "get")
	if added {
		s.logger.Info("权限添加成功")
	}
	if err := req.Validate(); err != nil {
		return nil, errorsx.ValidateError(errorsx.WithMessage(err.Error()))
	}

	return &HelloResponse{
		Msg: fmt.Sprintf("Hello, %s", req.Name),
	}, nil
}
