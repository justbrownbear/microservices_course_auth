package tests

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

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
	defer test.Cleanup(mc.Finish)

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

				return user_service.New(userRepositoryMock)
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

				return user_service.New(userRepositoryMock)
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

				return user_service.New(userRepositoryMock)
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

