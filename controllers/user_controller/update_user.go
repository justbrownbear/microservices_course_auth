package user_controller

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	grpc_converter "github.com/justbrownbear/microservices_course_auth/controllers/user_controller/converter"
	"github.com/justbrownbear/microservices_course_auth/pkg/user_v1"
)

func (instance *controller) UpdateUser(
	ctx context.Context,
	request *user_v1.UpdateUserRequest,
) (*emptypb.Empty, error) {
	payload := grpc_converter.ConvertUpdateUserFromGrpcModel(request)

	result := &emptypb.Empty{}

	err := instance.grpcAPI.UpdateUser(ctx, payload)
	if err != nil {
		return result, err
	}

	return result, nil
}
