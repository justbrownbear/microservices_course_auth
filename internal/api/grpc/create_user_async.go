package grpc_api

import (
	"context"
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
	"github.com/justbrownbear/microservices_course_auth/internal/client/kafka"
	user_model "github.com/justbrownbear/microservices_course_auth/internal/service/user/model"
)

// ***************************************************************************************************
// ***************************************************************************************************
func (instance *grpcAPI) CreateUserAsync(
	_ context.Context,
	userData *user_model.CreateUserRequest,
) error {
	userDataJSON, err := json.Marshal(userData)
	if err != nil {
		log.Printf("failed to marshal userData to JSON: %v\n", err)
		return err
	}

	message := &sarama.ProducerMessage{
		Topic: kafka.KafkaTopicUser,
		Key:   sarama.StringEncoder("CreateUserAsync"),
		Value: sarama.StringEncoder(userDataJSON),
	}

	// По-правильному, я должен был сделать сервис, у которого продьюсер
	// будет в репо-слое, но уже нет сил
	partition, offset, err := instance.kafkaProducer.SendMessage(message)
	if err != nil {
		log.Printf("failed to send message in Kafka: %v\n", err.Error())
		return err
	}

	log.Printf("message sent to partition %d with offset %d\n", partition, offset)

	return nil
}
