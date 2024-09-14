package tests

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	redis_client_mock "github.com/justbrownbear/microservices_course_auth/internal/client/cache/mocks"
	user_repository_mock "github.com/justbrownbear/microservices_course_auth/internal/repository/user/mocks"
	user_service "github.com/justbrownbear/microservices_course_auth/internal/service/user"
	user_model "github.com/justbrownbear/microservices_course_auth/internal/service/user/model"
	"github.com/justbrownbear/microservices_course_auth/internal/service/user/user_converter"
	"github.com/stretchr/testify/require"
)

func TestUpdateUser(test *testing.T) {
	test.Parallel()

	type args struct {
		ctx      context.Context
		userData *user_model.UpdateUserRequest
	}

	mc := minimock.NewController(test)

	ctx := context.Background()
	userID := gofakeit.Uint64()
	name := gofakeit.Name()
	email := gofakeit.Email()
	role := user_model.Role(1) // User

	updateUserRequest := &user_model.UpdateUserRequest{
		ID:    userID,
		Name:  name,
		Email: email,
		Role:  role,
	}

	serviceError := fmt.Errorf("service error")

	tests := []struct {
		name string
		args args
		err  error
		mock func(mc *minimock.Controller) user_service.UserService
	}{
		{
			name: "success case",
			args: args{
				ctx:      ctx,
				userData: updateUserRequest,
			},
			err: nil,
			mock: func(mc *minimock.Controller) user_service.UserService {
				userRepositoryMock := user_repository_mock.NewUserRepositoryMock(mc)
				userRepositoryMock.UpdateUserMock.Expect(ctx, user_converter.UpdateUserConvertRequest(updateUserRequest)).Return(nil)

				cacheMock := redis_client_mock.NewRedisClientMock(mc)
				cacheMock.DelMock.Return(nil)

				return user_service.New(userRepositoryMock, cacheMock)
			},
		},
		{
			name: "fail case",
			args: args{
				ctx:      ctx,
				userData: updateUserRequest,
			},
			err: serviceError,
			mock: func(mc *minimock.Controller) user_service.UserService {
				userRepositoryMock := user_repository_mock.NewUserRepositoryMock(mc)
				userRepositoryMock.UpdateUserMock.Expect(ctx, user_converter.UpdateUserConvertRequest(updateUserRequest)).Return(serviceError)

				cacheMock := redis_client_mock.NewRedisClientMock(mc)

				return user_service.New(userRepositoryMock, cacheMock)
			},
		},
		{
			name: "invalid userID case",
			args: args{
				ctx: ctx,
				userData: &user_model.UpdateUserRequest{
					ID:    0,
					Name:  name,
					Email: email,
					Role:  role,
				},
			},
			err: errors.New("user ID is required"),
			mock: func(mc *minimock.Controller) user_service.UserService {
				userRepositoryMock := user_repository_mock.NewUserRepositoryMock(mc)
				cacheMock := redis_client_mock.NewRedisClientMock(mc)

				return user_service.New(userRepositoryMock, cacheMock)
			},
		},
		{
			name: "invalid email format case",
			args: args{
				ctx: ctx,
				userData: &user_model.UpdateUserRequest{
					ID:    userID,
					Name:  name,
					Email: "invalid@email",
					Role:  role,
				},
			},
			err: errors.New("email has invalid format"),
			mock: func(mc *minimock.Controller) user_service.UserService {
				userRepositoryMock := user_repository_mock.NewUserRepositoryMock(mc)
				cacheMock := redis_client_mock.NewRedisClientMock(mc)

				return user_service.New(userRepositoryMock, cacheMock)
			},
		},
	}

	for _, testCase := range tests {
		testCase := testCase
		test.Run(testCase.name, func(t *testing.T) {
			userServiceMock := testCase.mock(mc)

			err := userServiceMock.UpdateUser(testCase.args.ctx, testCase.args.userData)
			require.Equal(t, testCase.err, err)
		})
	}
}
