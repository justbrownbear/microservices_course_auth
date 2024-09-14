package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	grpc_api "github.com/justbrownbear/microservices_course_auth/internal/api/grpc"
	user_service_mock "github.com/justbrownbear/microservices_course_auth/internal/service/user/mocks"
	user_model "github.com/justbrownbear/microservices_course_auth/internal/service/user/model"
	service_provider_mock "github.com/justbrownbear/microservices_course_auth/internal/service_provider/mocks"
	"github.com/justbrownbear/microservices_course_auth/internal/transaction_manager"
	transaction_manager_mock "github.com/justbrownbear/microservices_course_auth/internal/transaction_manager/mocks"
)

func TestUpdateUser(test *testing.T) {
	test.Parallel()

	// Создаем структуру входных параметров
	type args struct {
		ctx      context.Context
		userData *user_model.UpdateUserRequest
	}

	mc := minimock.NewController(test)

	// Делаем залипухи
	ctx := context.Background()
	userID := gofakeit.Uint64()
	name := gofakeit.Name()
	email := gofakeit.Email()
	role := user_model.Role(1) // User

	request := &user_model.UpdateUserRequest{
		ID:    userID,
		Name:  name,
		Email: email,
		Role:  role,
	}
	serviceError := fmt.Errorf("service error")

	// Объявим тип для функции, которая будет возвращать моки сервисов
	type grpcAPIMockFunction func(mc *minimock.Controller) grpc_api.GrpcAPI

	// Создаем набор тестовых кейсов
	tests := []struct {
		name string
		args args
		err  error
		mock grpcAPIMockFunction
	}{
		{
			name: "success case",
			args: args{
				ctx:      ctx,
				userData: request,
			},
			err: nil,
			mock: func(mc *minimock.Controller) grpc_api.GrpcAPI {
				// Делаем мок TxManager
				txManagerMock := transaction_manager_mock.NewTxManagerMock(mc)
				txManagerMock.WithTransactionMock.Set(
					func(ctx context.Context, handler transaction_manager.Handler) error {
						serviceProviderMock := service_provider_mock.NewServiceProviderMock(mc)
						userServiceMock := user_service_mock.NewUserServiceMock(mc)
						userServiceMock.UpdateUserMock.Expect(ctx, request).Return(nil)

						serviceProviderMock.GetUserServiceMock.Return(userServiceMock)

						return handler(ctx, serviceProviderMock)
					},
				)

				// Инициализируем GrpcAPI моком TxManager и ChatService
				return grpc_api.InitGrpcAPI(txManagerMock)
			},
		},
		{
			name: "fail case",
			args: args{
				ctx:      ctx,
				userData: request,
			},
			err: serviceError,
			mock: func(mc *minimock.Controller) grpc_api.GrpcAPI {
				// Делаем мок TxManager
				txManagerMock := transaction_manager_mock.NewTxManagerMock(mc)
				txManagerMock.WithTransactionMock.Set(
					func(ctx context.Context, handler transaction_manager.Handler) error {
						serviceProviderMock := service_provider_mock.NewServiceProviderMock(mc)
						userServiceMock := user_service_mock.NewUserServiceMock(mc)
						userServiceMock.UpdateUserMock.Expect(ctx, request).Return(serviceError)

						serviceProviderMock.GetUserServiceMock.Return(userServiceMock)

						return handler(ctx, serviceProviderMock)
					},
				)

				// Инициализируем GrpcAPI моком TxManager и ChatService
				return grpc_api.InitGrpcAPI(txManagerMock)
			},
		},
	}

	for _, testCase := range tests {
		testCase := testCase
		test.Run(testCase.name, func(t *testing.T) {
			grpcAPIMock := testCase.mock(mc)

			err := grpcAPIMock.UpdateUser(testCase.args.ctx, testCase.args.userData)
			require.Equal(t, testCase.err, err)
		})
	}
}
