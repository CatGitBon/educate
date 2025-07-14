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

func (a Auth) CreateUser() {
	// TODO: Implement createUser method
	a.logger.Info("createUser method called")
}

func (a Auth) DeleteUser() {
	// TODO: Implement deleteUser method
	a.logger.Info("deleteUser method called")
}

func (a Auth) UpdateUser() {
	// TODO: Implement updateUser method
	a.logger.Info("updateUser method called")
}

func (a Auth) GetUserByID() {
	// TODO: Implement getUserByID method
	a.logger.Info("getUserByID method called")
}
