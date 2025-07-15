package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/vctrl/currency-service/gateway/internal/dto"
	"github.com/vctrl/currency-service/gateway/internal/repository"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type authClientInterface interface {
	GenerateToken(ctx context.Context, login string) (string, error)
	ValidateToken(ctx context.Context, token string) error
}

type AuthService struct {
	authClient authClientInterface
	userRepo   repository.UserRepository // todo interface
}

func NewAuth(authClient authClientInterface, userRepo repository.UserRepository) AuthService {
	return AuthService{
		authClient: authClient,
		userRepo:   userRepo,
	}
}

func (s *AuthService) Register(req dto.RegisterRequest) error {

	// s.authClient.
	// TODO: Реализовать регистрацию через auth сервис
	// Пока просто возвращаем успех
	return nil
}

func (s *AuthService) Login(ctx context.Context, login, password string) (string, error) {
	// Проверяем, что пользователь существует
	_, err := s.userRepo.GetUser(ctx, 11)
	if err != nil {
		return "", fmt.Errorf("userRepo.GetUser: %w", err)
	}

	// Генерируем токен (проверка пароля будет в auth сервисе)
	res, err := s.authClient.GenerateToken(ctx, login)
	if err != nil {
		return "", fmt.Errorf("authClient.GenerateToken: %w", err)
	}

	return res, nil
}

func (s *AuthService) ValidateToken(ctx context.Context, token string) error {
	err := s.authClient.ValidateToken(ctx, token)
	if err != nil {
		return fmt.Errorf("authClient.ValidateToken: %w", err)
	}

	return nil
}

func (s *AuthService) Logout(token string) error {
	return errors.New("logout is not implemented")
}

func (s *AuthService) GetUserById(ctx context.Context, id int64) (repository.User, error) {
	return s.userRepo.GetUser(ctx, id)
}
