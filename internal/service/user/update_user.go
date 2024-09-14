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
func (instance *userService) UpdateUser(ctx context.Context, userData *user_model.UpdateUserRequest) error {
	err := updateUserValidateInputData(userData)
	if err != nil {
		return err
	}

	err = instance.updateUserWithCache(ctx, userData)
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


// ***************************************************************************************************
// ***************************************************************************************************
func (instance *userService) updateUserWithCache(
	ctx context.Context,
	userData *user_model.UpdateUserRequest,
) error {
	err := instance.updateUser(ctx, userData)
	if err != nil {
		return err
	}

	// Очищаем кэш пользователя
	cacheKey := getCacheKey( userData.ID )
	err = instance.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}

	return nil
}


// ***************************************************************************************************
// ***************************************************************************************************
func (instance *userService) updateUser(
	ctx context.Context,
	userData *user_model.UpdateUserRequest,
) error {
	payload := user_converter.UpdateUserConvertRequest(userData)

	err := instance.repository.UpdateUser(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}
