package user_controller

import (
	"context"
	"log"

	user_model "github.com/justbrownbear/microservices_course_auth/internal/service/user/model"
	"github.com/justbrownbear/microservices_course_auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ***************************************************************************************************
// ***************************************************************************************************
func (instance *controller) CreateUserAsync(
	ctx context.Context,
	request *user_v1.CreateUserRequest,
) (*emptypb.Empty, error) {
	payload := &user_model.CreateUserRequest{
		Name:            request.Name,
		Email:           request.Email,
		Password:        request.Password,
		PasswordConfirm: request.PasswordConfirm,
		Role:            user_model.Role(request.Role),
	}

	err := instance.grpcAPI.CreateUserAsync(ctx, payload)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	result := &emptypb.Empty{}

	return result, nil
}
