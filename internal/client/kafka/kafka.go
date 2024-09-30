package kafka

import (
	"context"

	kafka_consumer "github.com/justbrownbear/microservices_course_auth/internal/client/kafka/consumer"
)

// KafkaTopicUser - Топик Kafka для сообщений управления пользователями
var KafkaTopicUser = "user"

// Consumer - Интерфейс консьюмера Kafka
type Consumer interface {
	Consume(ctx context.Context, topicName string, handler kafka_consumer.Handler) error
	Close() error
}
