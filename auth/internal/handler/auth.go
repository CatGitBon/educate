package handler

import (
	"context"

	"github.com/vctrl/currency-service/auth/internal/dto"
	"github.com/vctrl/currency-service/pkg/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *AuthServer) GetUserById(ctx context.Context, req *auth.GetUserByIdRequest) (*auth.GetUserByIdResponse, error) {

	dtoReq := dto.GetUserByIdToDto(req)

	user := s.service.GetUserByID(dtoReq)

	if user == nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	return dto.GetUserByIdToProtoBuf(user), nil
}
