package grpc_api

import (
	"context"

	"github.com/IBM/sarama"
	user_model "github.com/justbrownbear/microservices_course_auth/internal/service/user/model"
	"github.com/justbrownbear/microservices_course_auth/internal/transaction_manager"
)

// GrpcAPI - Интерфейс gRPC API
type GrpcAPI interface {
	// Создание пользователя
	CreateUser(ctx context.Context, userData *user_model.CreateUserRequest) (uint64, error)
	// Создание пользователя асинхронно (через Kafka)
	CreateUserAsync(ctx context.Context, userData *user_model.CreateUserRequest) error
	// Получение данных пользователя по id
	GetUser(ctx context.Context, userID uint64) (*user_model.GetUserResponse, error)
	// Обновление данных пользователя
	UpdateUser(ctx context.Context, userData *user_model.UpdateUserRequest) error
	// Удаление пользователя
	DeleteUser(ctx context.Context, userID uint64) error
}

type grpcAPI struct {
	txManager     transaction_manager.TxManager
	kafkaProducer sarama.SyncProducer
}

// InitGrpcAPI инициализирует gRPC API
func InitGrpcAPI(txManager transaction_manager.TxManager, kafkaProducer sarama.SyncProducer) GrpcAPI {
	return &grpcAPI{
		txManager:     txManager,
		kafkaProducer: kafkaProducer,
	}
}
