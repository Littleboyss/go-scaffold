package role

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go-scaffold/internal/app/model"
	errorsx "go-scaffold/internal/app/pkg/errors"
)

// CreateRequest 创建角色请求参数
type CreateRequest struct {
	Name   string `json:"name"`
	Remark string `json:"remark"`
}

func (r CreateRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required.Error("名称不能为空")),
	)
}

// CreateResponse 创建角色响应数据
type CreateResponse struct {
	Id     uint64 `json:"id"`
	Name   string `json:"name"`
	Remark string `json:"phone"`
}

// Create 创建角色
func (s *Service) Create(ctx context.Context, req CreateRequest) (*CreateResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, errorsx.ValidateError(errorsx.WithMessage(err.Error()))
	}

	m := &model.Role{
		Name:   req.Name,
		Remark: req.Remark,
	}

	if _, err := s.repo.Create(ctx, m); err != nil {
		s.logger.Errorf("%s: %s", model.ErrDataStoreFailed, err)
		return nil, errorsx.ServerError(errorsx.WithMessage(model.ErrDataStoreFailed.Error()))
	}

	resp := &CreateResponse{
		Id:     m.Id,
		Name:   m.Name,
		Remark: m.Remark,
	}

	return resp, nil
}
