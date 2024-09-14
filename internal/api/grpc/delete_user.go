package grpc_api

import (
	"context"

	user_service "github.com/justbrownbear/microservices_course_auth/internal/service/user"
	"github.com/justbrownbear/microservices_course_auth/internal/service_provider"
)

// ***************************************************************************************************
// ***************************************************************************************************
func (instance *grpcAPI) DeleteUser(
	ctx context.Context,
	userID uint64,
) error {
	err := instance.txManager.WithTransaction(ctx,
		func(ctx context.Context, serviceProvider service_provider.ServiceProvider) error {
			// В этом месте нам пришел сервис-провайдер, который уже имеет connection внутри себя
			// Нам осталось только получить нужные сервисы, и...
			userService := serviceProvider.GetUserService()

			// ...передать их функции, которая на входе принимает только используемые сервисы и in
			err := deleteUser(ctx, userService, userID)
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
func deleteUser(
	ctx context.Context,
	userService user_service.UserService,
	userID uint64,
) error {
	err := userService.DeleteUser(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}
