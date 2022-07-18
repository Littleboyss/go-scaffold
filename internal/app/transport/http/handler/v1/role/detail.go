package role

import (
	"github.com/gin-gonic/gin"
	errorsx "go-scaffold/internal/app/pkg/errors"
	"go-scaffold/internal/app/service/role"
	"go-scaffold/internal/app/transport/http/pkg/response"
)

// Detail 角色详情
// @Router       /v1/role/{id} [get]
// @Summary      角色详情
// @Description  角色详情
// @Tags         角色
// @Accept       plain
// @Produce      json
// @Param        id   path      integer                                    true  "角色 id"  format(uint)  minimum(1)
// @Success      200  {object}  example.Success{data=role.DetailResponse}  "成功响应"
// @Failure      500  {object}  example.ServerError                        "服务器出错"
// @Failure      400  {object}  example.ClientError                        "客户端请求错误（code 类型应为 int，string 仅为了表达多个错误码）"
// @Failure      401  {object}  example.Unauthorized                       "登陆失效"
// @Failure      403  {object}  example.PermissionDenied                   "没有权限"
// @Failure      404  {object}  example.ResourceNotFound                   "资源不存在"
// @Failure      429  {object}  example.TooManyRequest                     "请求过于频繁"
// @Security     Authorization
func (h *Handler) Detail(ctx *gin.Context) {
	req := new(role.DetailRequest)
	if err := ctx.ShouldBindUri(req); err != nil {
		h.logger.Error(err)
		response.Error(ctx, errorsx.ValidateError())
		return
	}

	ret, err := h.service.Detail(ctx.Request.Context(), *req)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, response.WithData(ret))
	return
}
