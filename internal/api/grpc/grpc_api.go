package grpc_api

import (
	"context"

	user_model "github.com/justbrownbear/microservices_course_auth/internal/service/user/model"
	"github.com/justbrownbear/microservices_course_auth/internal/transaction_manager"
)

// GrpcAPI - Интерфейс gRPC API
type GrpcAPI interface {
	// Создание пользователя
	CreateUser(ctx context.Context, userData *user_model.CreateUserRequest) (uint64, error)
	// Получение данных пользователя по id
	GetUser(ctx context.Context, userID uint64) (*user_model.GetUserResponse, error)
	// Обновление данных пользователя
	UpdateUser(ctx context.Context, userData *user_model.UpdateUserRequest) error
	// Удаление пользователя
	DeleteUser(ctx context.Context, userID uint64) error
}


type grpcAPI struct {
	txManager transaction_manager.TxManager
}


// InitGrpcAPI инициализирует gRPC API
func InitGrpcAPI(txManager transaction_manager.TxManager) GrpcAPI {
	return &grpcAPI{
		txManager: txManager,
	}
}
