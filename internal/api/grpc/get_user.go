package grpc_api

import (
	"context"

	user_service "github.com/justbrownbear/microservices_course_auth/internal/service/user"
	user_model "github.com/justbrownbear/microservices_course_auth/internal/service/user/model"
	"github.com/justbrownbear/microservices_course_auth/internal/service_provider"
)

// ***************************************************************************************************
// ***************************************************************************************************
func (instance *grpcAPI) GetUser(
	ctx context.Context,
	userID uint64,
) (*user_model.GetUserResponse, error) {
	var result *user_model.GetUserResponse

	// TODO: Переделать на WithConnection
	err := instance.txManager.WithTransaction( ctx,
		func ( ctx context.Context, serviceProvider service_provider.ServiceProvider ) error {
			// В этом месте нам пришел сервис-провайдер, который уже имеет connection внутри себя
			// Нам осталось только получить нужные сервисы, и...
			userService := serviceProvider.GetUserService()

			// ...передать их функции, которая на входе принимает только используемые сервисы и in
			var err error
			result, err = getUser( ctx, userService, userID )
			if err != nil {
				return err
			}

			return nil
		} )
	if err != nil {
		return nil, err
	}

	return result, nil
}


// ***************************************************************************************************
// ***************************************************************************************************
func getUser(
	ctx context.Context,
	userService user_service.UserService,
	userID uint64,
) (*user_model.GetUserResponse, error) {
	result, err := userService.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
