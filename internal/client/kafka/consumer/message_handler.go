package kafka_consumer

import (
	"context"
	"log"

	"github.com/IBM/sarama"
)

// Handler - обработчик сообщения
type Handler func(ctx context.Context, message *sarama.ConsumerMessage) error

// GroupHandler - обработчик группы консюьмеров
type GroupHandler struct {
	messageHandler Handler
}

// NewGroupHandler - создает новый экземпляр GroupHandler
func NewGroupHandler() *GroupHandler {
	return &GroupHandler{}
}

// Setup - вызывается перед началом обработки сообщений
func (instance *GroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup - вызывается при завершенни
func (instance *GroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim должен запустить потребительский цикл сообщений ConsumerGroupClaim().
// После закрытия канала Messages() обработчик должен завершить обработку
func (instance *GroupHandler) ConsumeClaim(
	session sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim,
) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Println("ConsumerGroupClaim.Messages() channel closed")
				return nil
			}

			log.Printf("A message claimed: value = %s, timestamp = %v, topic = %s\n",
				string(message.Value),
				message.Timestamp,
				message.Topic)

			err := instance.messageHandler(session.Context(), message)
			if err != nil {
				log.Printf("Kafka message handler returned error: %s\n", err)
				continue
			}

			session.MarkMessage(message, "")

		case <-session.Context().Done():
			log.Println("ConsumerGroupSession.Context() done")
			return nil
		}
	}
}
