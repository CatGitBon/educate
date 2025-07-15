package handler

import (
	"context"

	"github.com/vctrl/currency-service/pkg/auth"
)

func (s *AuthServer) GetUserById(ctx context.Context, req *auth.GetUserByIdRequest) (*auth.GetUserByIdResponse, error) {
	// TODO: Implement actual user retrieval logic
	return &auth.GetUserByIdResponse{
		UserId:   req.UserId,
		Username: "test_user",
		Email:    "test@example.com",
	}, nil
}
