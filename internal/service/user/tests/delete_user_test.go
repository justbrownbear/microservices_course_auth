package tests

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	redis_client_mock "github.com/justbrownbear/microservices_course_auth/internal/client/cache/mocks"
	user_repository_mock "github.com/justbrownbear/microservices_course_auth/internal/repository/user/mocks"
	user_service "github.com/justbrownbear/microservices_course_auth/internal/service/user"
)

func TestDeleteUser(test *testing.T) {
	test.Parallel()

	type args struct {
		ctx    context.Context
		userID uint64
	}

	mc := minimock.NewController(test)

	ctx := context.Background()
	userID := gofakeit.Uint64()
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
				ctx:    ctx,
				userID: userID,
			},
			err: nil,
			mock: func(mc *minimock.Controller) user_service.UserService {
				userRepositoryMock := user_repository_mock.NewUserRepositoryMock(mc)
				userRepositoryMock.DeleteUserMock.Expect(ctx, int64(userID)).Return(nil)

				cacheMock := redis_client_mock.NewRedisClientMock(mc)
				cacheMock.DelMock.Return(nil)

				return user_service.New(userRepositoryMock, cacheMock)
			},
		},
		{
			name: "fail case",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			err: serviceError,
			mock: func(mc *minimock.Controller) user_service.UserService {
				userRepositoryMock := user_repository_mock.NewUserRepositoryMock(mc)
				userRepositoryMock.DeleteUserMock.Expect(ctx, int64(userID)).Return(serviceError)

				cacheMock := redis_client_mock.NewRedisClientMock(mc)

				return user_service.New(userRepositoryMock, cacheMock)
			},
		},
		{
			name: "invalid userID case",
			args: args{
				ctx:    ctx,
				userID: 0,
			},
			err: errors.New("user ID is required"),
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

			err := userServiceMock.DeleteUser(testCase.args.ctx, testCase.args.userID)
			require.Equal(t, testCase.err, err)
		})
	}
}
