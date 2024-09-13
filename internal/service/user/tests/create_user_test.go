package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	user_repository_mock "github.com/justbrownbear/microservices_course_auth/internal/repository/user/mocks"
	user_service "github.com/justbrownbear/microservices_course_auth/internal/service/user"
	user_model "github.com/justbrownbear/microservices_course_auth/internal/service/user/model"
)

func TestCreateUser(test *testing.T) {
	test.Parallel()

	// Создаем структуру входных параметров
	type args struct {
		ctx context.Context
		userData *user_model.CreateUserRequest
	}

	mc := minimock.NewController(test)
	defer test.Cleanup(mc.Finish)

	// Делаем залипухи
	ctx			:= context.Background()
	userID		:= gofakeit.Uint64()
	name		:= gofakeit.Name()
	email		:= gofakeit.Email()
	password	:= gofakeit.Password(true, true, true, true, true, 10)
	passwordConfirm := password
	role		:= user_model.Role( 1 ) // User

	request		:= &user_model.CreateUserRequest {
		Name: name,
		Email: email,
		Password: password,
		PasswordConfirm: passwordConfirm,
		Role: role,
	}
	response	:= userID
	serviceError := fmt.Errorf("service error")

	// Создаем набор тестовых кейсов
	tests := []struct {
		name	string
		args	args
		want	uint64
		err		error
		mock	func(mc *minimock.Controller) user_service.UserService
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				userData: request,
			},
			want: response,
			err: nil,
			mock: func(mc *minimock.Controller) user_service.UserService {
				userRespositoryMock := user_repository_mock.NewUserRepositoryMock(mc)
				userRespositoryMock.CreateUserMock.Return(int64(response), nil)

				return user_service.New(userRespositoryMock)
			},
		},
		{
			name: "fail case",
			args: args{
				ctx: ctx,
				userData: request,
			},
			want: 0,
			err: serviceError,
			mock: func(mc *minimock.Controller) user_service.UserService {
				userRespositoryMock := user_repository_mock.NewUserRepositoryMock(mc)
				userRespositoryMock.CreateUserMock.Return(0, serviceError)

				return user_service.New(userRespositoryMock)
			},
		},
		{
			name: "name is required fail case 1",
			args: args{
				ctx: ctx,
				userData: func() *user_model.CreateUserRequest {
                    copyRequest := *request
                    copyRequest.Name = ""
                    return &copyRequest
                }(),
			},
			want: 0,
			err: fmt.Errorf("name is required"),
			mock: func(mc *minimock.Controller) user_service.UserService {
				userRespositoryMock := user_repository_mock.NewUserRepositoryMock(mc)

				return user_service.New(userRespositoryMock)
			},
		},
		{
			name: "name is required fail case 2",
			args: args{
				ctx: ctx,
				userData: func() *user_model.CreateUserRequest {
                    copyRequest := *request
                    copyRequest.Name = " "
                    return &copyRequest
                }(),
			},
			want: 0,
			err: fmt.Errorf("name is required"),
			mock: func(mc *minimock.Controller) user_service.UserService {
				userRespositoryMock := user_repository_mock.NewUserRepositoryMock(mc)

				return user_service.New(userRespositoryMock)
			},
		},
		{
			name: "email is required fail case 1",
			args: args{
				ctx: ctx,
				userData: func() *user_model.CreateUserRequest {
                    copyRequest := *request
                    copyRequest.Email = ""
                    return &copyRequest
                }(),
			},
			want: 0,
			err: fmt.Errorf("email is required"),
			mock: func(mc *minimock.Controller) user_service.UserService {
				userRespositoryMock := user_repository_mock.NewUserRepositoryMock(mc)

				return user_service.New(userRespositoryMock)
			},
		},
		{
			name: "email is required fail case 2",
			args: args{
				ctx: ctx,
				userData: func() *user_model.CreateUserRequest {
                    copyRequest := *request
                    copyRequest.Email = " "
                    return &copyRequest
                }(),
			},
			want: 0,
			err: fmt.Errorf("email is required"),
			mock: func(mc *minimock.Controller) user_service.UserService {
				userRespositoryMock := user_repository_mock.NewUserRepositoryMock(mc)

				return user_service.New(userRespositoryMock)
			},
		},
		{
			name: "email has invalid format fail case 1",
			args: args{
				ctx: ctx,
				userData: func() *user_model.CreateUserRequest {
                    copyRequest := *request
                    copyRequest.Email = "fine"
                    return &copyRequest
                }(),
			},
			want: 0,
			err: fmt.Errorf("email has invalid format"),
			mock: func(mc *minimock.Controller) user_service.UserService {
				userRespositoryMock := user_repository_mock.NewUserRepositoryMock(mc)

				return user_service.New(userRespositoryMock)
			},
		},
		{
			name: "email has invalid format fail case 2",
			args: args{
				ctx: ctx,
				userData: func() *user_model.CreateUserRequest {
                    copyRequest := *request
                    copyRequest.Email = "fine@mail"
                    return &copyRequest
                }(),
			},
			want: 0,
			err: fmt.Errorf("email has invalid format"),
			mock: func(mc *minimock.Controller) user_service.UserService {
				userRespositoryMock := user_repository_mock.NewUserRepositoryMock(mc)

				return user_service.New(userRespositoryMock)
			},
		},
		{
			name: "password is required fail case",
			args: args{
				ctx: ctx,
				userData: func() *user_model.CreateUserRequest {
                    copyRequest := *request
                    copyRequest.Password = ""
                    return &copyRequest
                }(),
			},
			want: 0,
			err: fmt.Errorf("password is required"),
			mock: func(mc *minimock.Controller) user_service.UserService {
				userRespositoryMock := user_repository_mock.NewUserRepositoryMock(mc)

				return user_service.New(userRespositoryMock)
			},
		},
		{
			name: "password confirmation is required fail case",
			args: args{
				ctx: ctx,
				userData: func() *user_model.CreateUserRequest {
                    copyRequest := *request
                    copyRequest.PasswordConfirm = ""
                    return &copyRequest
                }(),
			},
			want: 0,
			err: fmt.Errorf("password confirmation is required"),
			mock: func(mc *minimock.Controller) user_service.UserService {
				userRespositoryMock := user_repository_mock.NewUserRepositoryMock(mc)

				return user_service.New(userRespositoryMock)
			},
		},
		{
			name: "passwords do not match fail case",
			args: args{
				ctx: ctx,
				userData: func() *user_model.CreateUserRequest {
                    copyRequest := *request
                    copyRequest.PasswordConfirm = copyRequest.Password + "!!!"
                    return &copyRequest
                }(),
			},
			want: 0,
			err: fmt.Errorf("passwords do not match"),
			mock: func(mc *minimock.Controller) user_service.UserService {
				userRespositoryMock := user_repository_mock.NewUserRepositoryMock(mc)

				return user_service.New(userRespositoryMock)
			},
		},
	}

	for _, testCase := range tests {
		testCase := testCase
		test.Run(testCase.name, func(t *testing.T) {
			userServiceMock := testCase.mock(mc)

			userID, err := userServiceMock.CreateUser(testCase.args.ctx, testCase.args.userData);
			require.Equal(t, testCase.err, err)
			require.Equal(t, testCase.want, userID)
		})
	}
}
