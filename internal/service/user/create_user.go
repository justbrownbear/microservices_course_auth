package user_service

import (
	"context"
	"errors"
	"strings"

	user_repository "github.com/justbrownbear/microservices_course_auth/internal/repository/user"
	"github.com/justbrownbear/microservices_course_auth/internal/service/password_management"
	user_model "github.com/justbrownbear/microservices_course_auth/internal/service/user/model"
	"github.com/justbrownbear/microservices_course_auth/internal/validator"
)

// ***************************************************************************************************
// ***************************************************************************************************
func (userServiceInstance *userService) CreateUser(
	ctx context.Context,
	userData *user_model.CreateUserRequest,
) (uint64, error) {
	err := createUserValidateInputData(userData)
	if err != nil {
		return 0, err
	}

	passwordHash, err := password_management.HashPassword(userData.Password)
	if err != nil {
		return 0, err
	}

	payload := user_repository.CreateUserParams{
		Name:         userData.Name,
		Email:        userData.Email,
		Role:         int16(userData.Role),
		PasswordHash: passwordHash,
	}

	userID, err := userServiceInstance.repository.CreateUser(ctx, payload)
	if err != nil {
		return 0, err
	}

	return uint64(userID), nil
}

// ***************************************************************************************************
// ***************************************************************************************************
func createUserValidateInputData(userData *user_model.CreateUserRequest) error {
	if len(strings.TrimSpace(userData.Name)) == 0 {
		return errors.New("name is required")
	}

	if len(strings.TrimSpace(userData.Email)) == 0 {
		return errors.New("email is required")
	}

	if !validator.IsValidEmail(userData.Email) {
		return errors.New("email has invalid format")
	}

	if len(strings.TrimSpace(userData.Password)) == 0 {
		return errors.New("password is required")
	}

	if len(strings.TrimSpace(userData.PasswordConfirm)) == 0 {
		return errors.New("password confirmation is required")
	}

	if userData.Password != userData.PasswordConfirm {
		return errors.New("passwords do not match")
	}

	return nil
}
