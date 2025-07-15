package handler

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/vctrl/currency-service/auth/internal/dto"
	"github.com/vctrl/currency-service/auth/internal/service"
	"github.com/vctrl/currency-service/pkg/auth"
	"go.uber.org/zap"
)

type AuthService interface {
	GetUserByID(dto *dto.GetUserByIdRequest) *dto.User
}

type AuthServer struct {
	auth.UnimplementedAuthServiceServer
	service AuthService
	logger  *zap.Logger

	requestCount    *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
	appUptime       prometheus.Gauge
}

func NewAuthServer(
	svc *service.Auth,
	logger *zap.Logger,
	requestCount *prometheus.CounterVec,
	requestDuration *prometheus.HistogramVec,
	appUptime prometheus.Gauge,
) *AuthServer {

	return &AuthServer{
		service:         svc,
		logger:          logger,
		requestCount:    requestCount,
		requestDuration: requestDuration,
		appUptime:       appUptime,
	}
}
