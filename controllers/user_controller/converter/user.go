package grpc_converter

import (
	user_model "github.com/justbrownbear/microservices_course_auth/internal/service/user/model"
	"github.com/justbrownbear/microservices_course_auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)


func ConvertGetUserFromDbModelToGrpcModel(
	userData *user_model.GetUserResponse,
) *user_v1.GetUserResponse {
	var updatedAt *timestamppb.Timestamp

	if userData.UpdatedAt.Valid {
		updatedAt = timestamppb.New(userData.UpdatedAt.Time)
	}

	result := &user_v1.GetUserResponse{
		Id: userData.ID,
		Name: userData.Name,
		Email: userData.Email,
		Role: user_v1.Role(userData.Role),
		CreatedAt: timestamppb.New(userData.CreatedAt),
		UpdatedAt: updatedAt,
	}

	return result
}


func ConvertUpdateUserFromGrpcModel(
	request *user_v1.UpdateUserRequest,
) *user_model.UpdateUserRequest {
	result := &user_model.UpdateUserRequest{
		ID:		request.GetId(),
		Name:	request.GetName(),
		Email:	request.GetEmail(),
		Role:	user_model.Role(request.GetRole()),
	}

	return result
}
