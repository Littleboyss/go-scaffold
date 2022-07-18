package role

import (
	"context"
	"go-scaffold/internal/app/model"
	errorsx "go-scaffold/internal/app/pkg/errors"
	"go-scaffold/internal/app/repository/role"
)

// ListRequest 角色列表请求参数
type ListRequest struct {
	Keyword string `json:"keyword" form:"keyword"`
}

// ListItem 角色列表项
type ListItem struct {
	Id     uint64 `json:"id"`
	Name   string `json:"name"`
	Remark string `json:"remark"`
}

// ListResponse 角色列表响应数据
type ListResponse []*ListItem

// List 角色列表
func (s *Service) List(ctx context.Context, req ListRequest) (ListResponse, error) {
	list, err := s.repo.FindList(
		ctx,
		role.FindListParam{Keyword: req.Keyword},
		[]string{"*"},
		"updated_at DESC",
	)
	if err != nil {
		s.logger.Errorf("%s: %s", model.ErrDataQueryFailed, err)
		return nil, errorsx.ServerError(errorsx.WithMessage(model.ErrDataQueryFailed.Error()))
	}

	resp := make(ListResponse, 0, len(list))

	for _, item := range list {
		resp = append(resp, &ListItem{
			Id:     item.Id,
			Name:   item.Name,
			Remark: item.Remark,
		})
	}

	return resp, nil
}
