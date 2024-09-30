package user_converter

import (
	"database/sql"
	"strconv"
	"time"

	user_repository "github.com/justbrownbear/microservices_course_auth/internal/repository/user"
	user_model "github.com/justbrownbear/microservices_course_auth/internal/service/user/model"
	"github.com/justbrownbear/microservices_course_auth/internal/type_converters"
)

// ToGetUserResponseFromRepository converts a GetUserRow from the repository layer
// to a GetUserResponse model used in the service layer.
//
// Parameters:
//   - repoModel: A pointer to a user_repository.GetUserRow containing user data from the repository.
//
// Returns:
//   - A pointer to a user_model.GetUserResponse containing the converted user data.
func ToGetUserResponseFromRepository(repoModel *user_repository.GetUserRow) *user_model.GetUserResponse {
	return &user_model.GetUserResponse{
		ID:        uint64(repoModel.ID),
		Name:      repoModel.Name,
		Email:     repoModel.Email,
		Role:      user_model.Role(repoModel.Role),
		CreatedAt: repoModel.CreateTimestamp.Time,
		UpdatedAt: sql.NullTime{Time: repoModel.UpdateTimestamp.Time, Valid: repoModel.UpdateTimestamp.Valid},
	}
}

// UpdateUserConvertRequest converts a user_model.UpdateUserRequest to user_repository.UpdateUserParams.
// It maps the fields from the input userData to the corresponding fields in the result.
//
// Parameters:
//   - userData: A pointer to user_model.UpdateUserRequest containing the user data to be converted.
//
// Returns:
//   - user_repository.UpdateUserParams: The converted user data suitable for repository operations.
func UpdateUserConvertRequest(userData *user_model.UpdateUserRequest) user_repository.UpdateUserParams {
	result := user_repository.UpdateUserParams{
		ID:    int64(userData.ID),
		Name:  userData.Name,
		Email: userData.Email,
		Role:  int16(userData.Role),
	}

	return result
}

// GetUserWithCacheConvertToRedis converts a GetUserResponse object to a map of strings
// suitable for caching in Redis. The map contains the user's ID, name, email, role,
// creation time, and update time.
//
// Parameters:
//   - userData: A pointer to a GetUserResponse object containing user data.
//
// Returns:
//
//	A map[string]string where the keys are field names and the values are the corresponding
//	string representations of the user's data.
func GetUserWithCacheConvertToRedis(userData *user_model.GetUserResponse) map[string]string {
	result := map[string]string{
		"id":         strconv.Itoa(int(userData.ID)),
		"name":       userData.Name,
		"email":      userData.Email,
		"role":       strconv.Itoa(int(userData.Role)),
		"created_at": type_converters.TimeToInt64String(userData.CreatedAt),
		"updated_at": type_converters.SQLNullTimeToInt64String(userData.UpdatedAt),
	}

	return result
}

// GetUserWithCacheConvertFromRedis converts user data from a Redis cache format
// (user_model.GetUserResponseForRedis) to a standard user response format
// (user_model.GetUserResponse).
//
// Parameters:
//   - userData: user_model.GetUserResponseForRedis
//     The user data retrieved from Redis cache.
//
// Returns:
//   - *user_model.GetUserResponse
//     The converted user data in the standard response format.
func GetUserWithCacheConvertFromRedis(userData user_model.GetUserResponseForRedis) *user_model.GetUserResponse {
	result := &user_model.GetUserResponse{
		ID:        userData.ID,
		Name:      userData.Name,
		Email:     userData.Email,
		Role:      userData.Role,
		CreatedAt: time.Unix(userData.CreatedAt, 0),
		UpdatedAt: type_converters.Int64ToSQLNullTime(userData.UpdatedAt),
	}

	return result
}
