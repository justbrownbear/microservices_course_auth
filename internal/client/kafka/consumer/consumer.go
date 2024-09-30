package kafka_consumer

import (
	"context"
	"errors"
	"log"

	"github.com/IBM/sarama"
)

type consumer struct {
	consumerGroup        sarama.ConsumerGroup
	consumerGroupHandler *GroupHandler
}

// NewConsumer создает новый экземпляр Consumer
func NewConsumer(
	consumerGroup sarama.ConsumerGroup,
	consumerGroupHandler *GroupHandler,
) *consumer {
	return &consumer{
		consumerGroup:        consumerGroup,
		consumerGroupHandler: consumerGroupHandler,
	}
}

// Consume запускает потребительский цикл сообщений
func (instance *consumer) Consume(
	ctx context.Context,
	topicName string,
	handler Handler,
) error {
	instance.consumerGroupHandler.messageHandler = handler

	return instance.consume(ctx, topicName)
}

// Close закрывает консьюмер
func (instance *consumer) Close() error {
	return instance.consumerGroup.Close()
}

// ***************************************************************************************************
// ***************************************************************************************************
func (instance *consumer) consume(
	ctx context.Context,
	topicName string,
) error {
	// Этот цикл нужен для того, чтобы после ребалансировки консьюмеров,
	// консьюмер смог снова подписаться на топики
	for {
		topicNameSlice := []string{topicName}
		err := instance.consumerGroup.Consume(ctx, topicNameSlice, instance.consumerGroupHandler)
		if err != nil {
			// Если консьюмер закрыт, то это не ошибка
			if errors.Is(err, sarama.ErrClosedConsumerGroup) {
				return nil
			}

			return err
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}

		log.Println("ConsumerGroup.Consume() returned. Probably rebalancing in progress. Reconnecting...")
	}
}
