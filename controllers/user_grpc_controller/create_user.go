package user_controller

import (
	"context"
	"log"

	user_model "github.com/justbrownbear/microservices_course_auth/internal/service/user/model"
	"github.com/justbrownbear/microservices_course_auth/pkg/user_v1"
)

// ***************************************************************************************************
// ***************************************************************************************************
func (controllerInstance *controller) CreateUser(
	ctx context.Context,
	request *user_v1.CreateUserRequest,
) (*user_v1.CreateUserResponse, error) {
	payload := &user_model.CreateUserRequest{
		Name:            request.Name,
		Email:           request.Email,
		Password:        request.Password,
		PasswordConfirm: request.PasswordConfirm,
		Role:            user_model.Role(request.Role),
	}

	userID, err := controllerInstance.grpcAPI.CreateUser(ctx, payload)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	result := &user_v1.CreateUserResponse{
		Id: userID,
	}

	return result, nil
}
