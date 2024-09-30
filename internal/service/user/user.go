package user_service

import (
	"context"

	"github.com/justbrownbear/microservices_course_auth/internal/client/cache"
	user_repository "github.com/justbrownbear/microservices_course_auth/internal/repository/user"
	user_model "github.com/justbrownbear/microservices_course_auth/internal/service/user/model"
)

// UserService defines the interface for user-related operations.
// It includes methods for creating, retrieving, updating, and deleting user data.
type UserService interface {
	// Создание пользователя
	CreateUser(ctx context.Context, userData *user_model.CreateUserRequest) (uint64, error)
	// Получение данных пользователя по id
	GetUser(ctx context.Context, userID uint64) (*user_model.GetUserResponse, error)
	// Обновление данных пользователя
	UpdateUser(ctx context.Context, userData *user_model.UpdateUserRequest) error
	// Удаление пользователя
	DeleteUser(ctx context.Context, userID uint64) error
}

type userService struct {
	repository user_repository.UserRepository
	cache      cache.RedisClient
}

// New инициализирует сервис пользователей
func New(repository user_repository.UserRepository, cache cache.RedisClient) UserService {
	return &userService{
		repository: repository,
		cache:      cache,
	}
}
