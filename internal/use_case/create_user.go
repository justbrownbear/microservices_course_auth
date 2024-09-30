package use_case

import (
	"context"

	user_service "github.com/justbrownbear/microservices_course_auth/internal/service/user"
	user_model "github.com/justbrownbear/microservices_course_auth/internal/service/user/model"
)

// CreateUser creates a new user with the provided user data.
func CreateUser(
	ctx context.Context,
	userService user_service.UserService,
	userData *user_model.CreateUserRequest,
) (uint64, error) {
	chatID, err := userService.CreateUser(ctx, userData)
	if err != nil {
		return 0, err
	}

	return chatID, nil
}
