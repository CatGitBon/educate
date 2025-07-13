package service

import (
	"github.com/vctrl/currency-service/auth/internal/repository"
	"go.uber.org/zap"
)

type Auth struct {
	currencyRepo repository.Auth
	logger       *zap.Logger
}

func NewAuth(
	repo repository.Auth,
	logger *zap.Logger,
) Auth {
	return Auth{
		currencyRepo: repo,
		logger:       logger,
	}
}

func (a Auth) GetUser() {
	// TODO: Implement getUser method
	a.logger.Info("getUser method called")
}
