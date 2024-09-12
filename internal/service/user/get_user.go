package user_service

import (
	"context"
	"errors"

	user_model "github.com/justbrownbear/microservices_course_auth/internal/service/user/model"
	"github.com/justbrownbear/microservices_course_auth/internal/service/user/user_converter"
)

// ***************************************************************************************************
// ***************************************************************************************************
func (userServiceInstance *userService) GetUser(
	ctx context.Context,
	userID uint64,
) (*user_model.GetUserResponse, error) {
	err := getUserValidateInputData(userID)
	if err != nil {
		return nil, err
	}

	userIdInt64 := int64(userID)

	userData, err := userServiceInstance.repository.GetUser(ctx, userIdInt64)
	if err != nil {
		return nil, err
	}

	result := user_converter.ToGetUserResponseFromRepository( &userData )

	return result, nil
}


// ***************************************************************************************************
// ***************************************************************************************************
func getUserValidateInputData( userID uint64 ) error {
	if userID <= 0 {
		return errors.New("user ID is required")
	}

	return nil
}
