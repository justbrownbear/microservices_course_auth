package user_service

import (
	"context"
	"errors"

	user_model "github.com/justbrownbear/microservices_course_auth/internal/service/user/model"
	"github.com/justbrownbear/microservices_course_auth/internal/service/user/user_converter"
	"github.com/justbrownbear/microservices_course_auth/internal/validator"
)

// ***************************************************************************************************
// ***************************************************************************************************
func (userServiceInstance *userService) UpdateUser(ctx context.Context, userData *user_model.UpdateUserRequest) error {
	err := updateUserValidateInputData(userData)
	if err != nil {
		return err
	}

	payload := user_converter.UpdateUserConvertRequest(userData)

	err = userServiceInstance.repository.UpdateUser(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}


// ***************************************************************************************************
// ***************************************************************************************************
func updateUserValidateInputData( userData *user_model.UpdateUserRequest ) error {
	if userData.ID <= 0 {
		return errors.New("user ID is required")
	}

	if len(userData.Email) > 0 && !validator.IsValidEmail(userData.Email) {
		return errors.New("email has invalid format")
	}

	return nil
}
