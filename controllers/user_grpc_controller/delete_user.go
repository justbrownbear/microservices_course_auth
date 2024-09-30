package user_controller

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/justbrownbear/microservices_course_auth/pkg/user_v1"
)

func (instance *controller) DeleteUser(
	ctx context.Context,
	request *user_v1.DeleteUserRequest,
) (*emptypb.Empty, error) {
	userID := request.GetId()

	err := instance.grpcAPI.DeleteUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := &emptypb.Empty{}

	return result, nil
}
