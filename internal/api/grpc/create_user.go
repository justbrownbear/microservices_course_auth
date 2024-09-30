package grpc_api

import (
	"context"

	user_model "github.com/justbrownbear/microservices_course_auth/internal/service/user/model"
	"github.com/justbrownbear/microservices_course_auth/internal/service_provider"
	"github.com/justbrownbear/microservices_course_auth/internal/use_case"
)

// ***************************************************************************************************
// ***************************************************************************************************
func (instance *grpcAPI) CreateUser(ctx context.Context, userData *user_model.CreateUserRequest) (uint64, error) {
	var userID uint64

	err := instance.txManager.WithTransaction(ctx,
		func(ctx context.Context, serviceProvider service_provider.ServiceProvider) error {
			// В этом месте нам пришел сервис-провайдер, который уже имеет connection внутри себя
			// Нам осталось только получить нужные сервисы, и...
			userService := serviceProvider.GetUserService()

			// ...передать их функции, которая на входе принимает только используемые сервисы и in
			var err error
			userID, err = use_case.CreateUser(ctx, userService, userData)
			if err != nil {
				return err
			}

			return nil
		})
	if err != nil {
		return 0, err
	}

	return userID, nil
}
