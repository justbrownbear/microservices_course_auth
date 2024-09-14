package grpc_api

import (
	"context"

	user_service "github.com/justbrownbear/microservices_course_auth/internal/service/user"
	user_model "github.com/justbrownbear/microservices_course_auth/internal/service/user/model"
	"github.com/justbrownbear/microservices_course_auth/internal/service_provider"
)

// ***************************************************************************************************
// ***************************************************************************************************
func (instance *grpcAPI) UpdateUser(
	ctx context.Context,
	userData *user_model.UpdateUserRequest,
) error {
	err := instance.txManager.WithTransaction(ctx,
		func(ctx context.Context, serviceProvider service_provider.ServiceProvider) error {
			// В этом месте нам пришел сервис-провайдер, который уже имеет connection внутри себя
			// Нам осталось только получить нужные сервисы, и...
			userService := serviceProvider.GetUserService()

			// ...передать их функции, которая на входе принимает только используемые сервисы и in
			err := updateUser(ctx, userService, userData)
			if err != nil {
				return err
			}

			return nil
		})
	if err != nil {
		return err
	}

	return nil
}

// ***************************************************************************************************
// ***************************************************************************************************
func updateUser(
	ctx context.Context,
	userService user_service.UserService,
	userData *user_model.UpdateUserRequest,
) error {
	err := userService.UpdateUser(ctx, userData)
	if err != nil {
		return err
	}

	return nil
}
