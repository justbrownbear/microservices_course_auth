package user_converter

import (
	user_repository "github.com/justbrownbear/microservices_course_auth/internal/repository/user"
	user_model "github.com/justbrownbear/microservices_course_auth/internal/service/user/model"
)


func ToGetUserResponseFromRepository( repoModel *user_repository.GetUserRow ) *user_model.GetUserResponse {
	return &user_model.GetUserResponse {
		ID: uint64(repoModel.ID),
		Name: repoModel.Name,
		Email: repoModel.Email,
		Role: user_model.Role(repoModel.Role),
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


// func nullableString( value string ) sql.NullString {
// 	isValid := len(value) != 0
// 	result := sql.NullString{String: value, Valid: isValid}

// 	return result;
// }
