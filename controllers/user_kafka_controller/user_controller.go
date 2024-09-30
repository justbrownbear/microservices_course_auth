package user_kafka_controller

import (
	"context"
	"log"

	"github.com/IBM/sarama"
	kafka_api "github.com/justbrownbear/microservices_course_auth/internal/api/kafka"
)

type controller struct {
	kafkaAPI kafka_api.KafkaAPI
}

// UserKafkaController - Интерфейс контроллера Kafka
type UserKafkaController interface {
	OnMessage(ctx context.Context, message *sarama.ConsumerMessage) error
}

// InitUserKafkaController инициализирует контроллер Kafka
func InitUserKafkaController(kafkaAPI kafka_api.KafkaAPI) *controller {
	return &controller{
		kafkaAPI: kafkaAPI,
	}
}

// OnMessage обрабатывает сообщение из Kafka
func (instance *controller) OnMessage(ctx context.Context, message *sarama.ConsumerMessage) error {
	key := string(message.Key)
	value := string(message.Value)

	log.Printf("Received message: key %s, value %s\n", key, value)

	if key == "CreateUserAsync" {
		user, err := instance.createUser(ctx, message)
		if err != nil {
			log.Printf("Failed to create user: %v\n", err)
			return err
		}

		log.Printf("User created with ID %d\n", user.Id)
	}

	return nil
}
