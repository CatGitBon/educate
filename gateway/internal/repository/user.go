package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/vctrl/currency-service/pkg/auth"
)

var (
	ErrUserAlreadyExist = errors.New("user already exist")
	ErrUserNotFound     = errors.New("user not found")
)

type User struct {
	Login    string
	Password string
}

type UserRepository struct {
	users      map[string]User
	mu         *sync.RWMutex
	authClient auth.AuthServiceClient
}

func NewUser(authClient auth.AuthServiceClient) UserRepository {
	return UserRepository{
		users:      make(map[string]User),
		mu:         &sync.RWMutex{},
		authClient: authClient,
	}
}

func (repo *UserRepository) AddUser(user User) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, exists := repo.users[user.Login]; exists {
		return ErrUserAlreadyExist
	}
	repo.users[user.Login] = user

	return nil
}

func (repo *UserRepository) GetUser(ctx context.Context, userId int64) (User, error) {
	// Делаем gRPC вызов к auth сервису
	req := &auth.GetUserByIdRequest{
		UserId: userId,
	}

	resp, err := repo.authClient.GetUserById(ctx, req)

	if err != nil {
		return User{}, ErrUserNotFound
	}

	// Преобразуем ответ от auth сервиса в локальную структуру User
	return User{
		Login: resp.Username,
		// Password не возвращается из auth сервиса по соображениям безопасности
	}, nil
}
