package kafka_api

import (
	"context"

	user_model "github.com/justbrownbear/microservices_course_auth/internal/service/user/model"
	"github.com/justbrownbear/microservices_course_auth/internal/transaction_manager"
)

// KafkaAPI - Интерфейс Kafka API
type KafkaAPI interface {
	// Создание пользователя
	CreateUser(ctx context.Context, userData *user_model.CreateUserRequest) (uint64, error)
}

type kafkaAPI struct {
	txManager transaction_manager.TxManager
}

// InitKafkaAPI инициализирует gRPC API
func InitKafkaAPI(txManager transaction_manager.TxManager) KafkaAPI {
	return &kafkaAPI{
		txManager: txManager,
	}
}
