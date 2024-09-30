package user_kafka_controller

import (
	"context"
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
	user_model "github.com/justbrownbear/microservices_course_auth/internal/service/user/model"
	"github.com/justbrownbear/microservices_course_auth/pkg/user_v1"
)

// ***************************************************************************************************
// ***************************************************************************************************
// CreateUser creates a new user with the provided user data.
func (instance *controller) createUser(
	ctx context.Context,
	message *sarama.ConsumerMessage,
) (*user_v1.CreateUserResponse, error) {
	value := string(message.Value)

	var payload user_model.CreateUserRequest
	if err := json.Unmarshal([]byte(value), &payload); err != nil {
		log.Printf("failed to unmarshal message value: %v\n", err)
		return nil, err
	}

	userID, err := instance.kafkaAPI.CreateUser(ctx, &payload)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	result := &user_v1.CreateUserResponse{
		Id: userID,
	}

	return result, nil
}
