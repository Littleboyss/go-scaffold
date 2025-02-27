package user

import (
	"context"
	"go-scaffold/internal/app/service/user"
	pb "go-scaffold/internal/app/transport/grpc/api/scaffold/v1/user"
)

func (h *Handler) Detail(ctx context.Context, req *pb.DetailRequest) (*pb.DetailResponse, error) {
	svcReq := user.DetailRequest{
		Id: req.Id,
	}

	ret, err := h.service.Detail(ctx, svcReq)
	if err != nil {
		return nil, err
	}

	resp := &pb.DetailResponse{
		Id:    ret.Id,
		Name:  ret.Name,
		Age:   int32(ret.Age),
		Phone: ret.Phone,
	}

	return resp, nil
}
