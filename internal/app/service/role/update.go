package role

import (
	"context"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go-scaffold/internal/app/model"
	errorsx "go-scaffold/internal/app/pkg/errors"
	"gorm.io/gorm"
)

// UpdateRequest 更新角色请求参数
type UpdateRequest struct {
	Id     uint64 `json:"id" uri:"id"`
	Name   string `json:"name"`
	Remark string `json:"remark"`
}

func (r UpdateRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Id, validation.Required.Error("id 不能为空")),
		validation.Field(&r.Name, validation.Required.Error("名称不能为空")),
	)
}

// UpdateResponse 更新角色响应数据
type UpdateResponse struct {
	Id     uint64 `json:"id"`
	Name   string `json:"name"`
	Remark string `json:"remark"`
}

// Update 更新角色
func (s *Service) Update(ctx context.Context, req UpdateRequest) (*UpdateResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, errorsx.ValidateError(errorsx.WithMessage(err.Error()))
	}

	m, err := s.repo.FindOneById(
		ctx,
		req.Id,
		[]string{"*"},
	)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorsx.ResourceNotFound(errorsx.WithMessage(model.ErrDataNotFound.Error()))
		}
		s.logger.Errorf("%s: %s", model.ErrDataQueryFailed, err)
		return nil, errorsx.ServerError(errorsx.WithMessage(model.ErrDataQueryFailed.Error()))
	}

	m = &model.Role{
		BaseModel: m.BaseModel,
		Name:      req.Name,
		Remark:    req.Remark,
	}

	if _, err = s.repo.Save(ctx, m); err != nil {
		s.logger.Errorf("%s: %s", model.ErrDataStoreFailed, err)
		return nil, errorsx.ServerError(errorsx.WithMessage(model.ErrDataStoreFailed.Error()))
	}

	resp := &UpdateResponse{
		Id:     m.Id,
		Name:   m.Name,
		Remark: m.Remark,
	}

	return resp, nil
}
