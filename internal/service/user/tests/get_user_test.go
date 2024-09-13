package tests

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	user_repository "github.com/justbrownbear/microservices_course_auth/internal/repository/user"
	user_repository_mock "github.com/justbrownbear/microservices_course_auth/internal/repository/user/mocks"
	user_service "github.com/justbrownbear/microservices_course_auth/internal/service/user"
	user_model "github.com/justbrownbear/microservices_course_auth/internal/service/user/model"
)


func TestGetUser(test *testing.T) {
	test.Parallel()

	type args struct {
		ctx    context.Context
		userID uint64
	}

	mc := minimock.NewController(test)
	defer test.Cleanup(mc.Finish)

	ctx := context.Background()
	userID := gofakeit.Uint64()
	name := gofakeit.Name()
	email := gofakeit.Email()
	role := user_model.Role(1) // User
	// createdAt := gofakeit.Date()
	// updatedAt := sql.NullTime{Time: gofakeit.Date(), Valid: true}

	getUserRepositoryResponse := user_repository.GetUserRow{
		ID: int64(userID),
		Name: name,
		Email: email,
		Role: int16(role),
	}

	response := &user_model.GetUserResponse{
		ID:        userID,
		Name:      name,
		Email:     email,
		Role:      role,
		// CreatedAt: createdAt,
		// UpdatedAt: updatedAt,
	}
	serviceError := fmt.Errorf("service error")

	tests := []struct {
		name string
		args args
		want *user_model.GetUserResponse
		err  error
		mock func(mc *minimock.Controller) user_service.UserService
	}{
		{
			name: "success case",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			want: response,
			err:  nil,
			mock: func(mc *minimock.Controller) user_service.UserService {
				userRepositoryMock := user_repository_mock.NewUserRepositoryMock(mc)
				userRepositoryMock.GetUserMock.Expect(ctx, int64(userID)).Return(getUserRepositoryResponse, nil)

				return user_service.New(userRepositoryMock)
			},
		},
		{
			name: "fail case",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			want: nil,
			err:  serviceError,
			mock: func(mc *minimock.Controller) user_service.UserService {
				userRepositoryMock := user_repository_mock.NewUserRepositoryMock(mc)
				userRepositoryMock.GetUserMock.Expect(ctx, int64(userID)).Return(user_repository.GetUserRow{}, serviceError)

				return user_service.New(userRepositoryMock)
			},
		},
		{
			name: "invalid userID case",
			args: args{
				ctx:    ctx,
				userID: 0,
			},
			want: nil,
			err:  errors.New("user ID is required"),
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

			result, err := userServiceMock.GetUser(testCase.args.ctx, testCase.args.userID)
			require.Equal(t, testCase.err, err)
			require.Equal(t, testCase.want, result)
		})
	}
}
