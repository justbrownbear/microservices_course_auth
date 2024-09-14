package user_converter

import (
	"database/sql"
	"strconv"
	"time"

	user_repository "github.com/justbrownbear/microservices_course_auth/internal/repository/user"
	user_model "github.com/justbrownbear/microservices_course_auth/internal/service/user/model"
	"github.com/justbrownbear/microservices_course_auth/internal/type_converters"
)


func ToGetUserResponseFromRepository( repoModel *user_repository.GetUserRow ) *user_model.GetUserResponse {
	return &user_model.GetUserResponse {
		ID: uint64(repoModel.ID),
		Name: repoModel.Name,
		Email: repoModel.Email,
		Role: user_model.Role(repoModel.Role),
		CreatedAt: repoModel.CreateTimestamp.Time,
		UpdatedAt: sql.NullTime{ Time: repoModel.UpdateTimestamp.Time, Valid: repoModel.UpdateTimestamp.Valid },
	}
}


func UpdateUserConvertRequest( userData *user_model.UpdateUserRequest ) user_repository.UpdateUserParams {
	result := user_repository.UpdateUserParams {
		ID: int64(userData.ID),
		Name: userData.Name,
		Email: userData.Email,
		Role: int16(userData.Role),
	}

	return result
}


func GetUserWithCacheConvertToRedis( userData *user_model.GetUserResponse ) map[string]string {
	result := map[string]string {
		"id": strconv.Itoa(int(userData.ID)),
		"name": userData.Name,
		"email": userData.Email,
		"role": strconv.Itoa(int(userData.Role)),
		"created_at": type_converters.TimeToInt64String( userData.CreatedAt ),
		"updated_at": type_converters.SqlNullTimeToInt64String( userData.UpdatedAt ),
	}

	return result
}


func GetUserWithCacheConvertFromRedis( userData user_model.GetUserResponseForRedis ) *user_model.GetUserResponse {
	result := &user_model.GetUserResponse {
		ID: userData.ID,
		Name: userData.Name,
		Email: userData.Email,
		Role: userData.Role,
		CreatedAt: time.Unix( userData.CreatedAt, 0 ),
		UpdatedAt: type_converters.Int64ToSqlNullTime( userData.UpdatedAt ),
	}

	return result
}

