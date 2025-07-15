package handler

import (
	"github.com/vctrl/currency-service/pkg/auth"
	"go.uber.org/zap"
	// "github.com/vctrl/currency-service/auth/internal/dto"
)

type AuthService interface {
	GetUserByID()
}

// todo tests
type AuthServer struct {
	auth.UnimplementedAuthServiceServer
	service AuthService
	logger  *zap.Logger

	// requestCount    *prometheus.CounterVec
	// requestDuration *prometheus.HistogramVec
	// appUptime       prometheus.Gauge
}

func NewAuthServer(svc AuthService,
	logger *zap.Logger,
	// requestCount *prometheus.CounterVec,
	// requestDuration *prometheus.HistogramVec,
	// appUptime prometheus.Gauge
) *AuthServer {

	return &AuthServer{
		service: svc,
		logger:  logger,
		// requestCount:    requestCount,
		// requestDuration: requestDuration,
		// appUptime:       appUptime,
	}
}

// func (s *AuthServer) GetUser(ctx context.Context, req *auth.GetUserRequest) (*auth.GetUserResponse, error) {
// 	// TODO: Implement actual user retrieval logic
// 	return &auth.GetUserResponse{
// 		UserId:   req.UserId,
// 		Username: "test_user",
// 		Email:    "test@example.com",
// 	}, nil
// }
