package user_controller

import (
	"context"
	"log"

	grpc_converter "github.com/justbrownbear/microservices_course_auth/controllers/user_grpc_controller/converter"
	"github.com/justbrownbear/microservices_course_auth/pkg/user_v1"
)

func (instance *controller) GetUser(
	ctx context.Context,
	request *user_v1.GetUserRequest,
) (*user_v1.GetUserResponse, error) {
	userID := request.GetId()

	userData, err := instance.grpcAPI.GetUser(ctx, userID)
	if err != nil {
		log.Printf("Failed to get user: %v", err)
		return nil, err
	}

	result := grpc_converter.ConvertGetUserFromDbModelToGrpcModel(userData)

	return result, nil
}
