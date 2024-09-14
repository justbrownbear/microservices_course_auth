package user_service

import (
	"context"
	"errors"

	"github.com/gomodule/redigo/redis"
	user_model "github.com/justbrownbear/microservices_course_auth/internal/service/user/model"
	"github.com/justbrownbear/microservices_course_auth/internal/service/user/user_converter"
)

// ***************************************************************************************************
// ***************************************************************************************************
func (userServiceInstance *userService) GetUser(
	ctx context.Context,
	userID uint64,
) (*user_model.GetUserResponse, error) {
	err := getUserValidateInputData(userID)
	if err != nil {
		return nil, err
	}

	// result, err := userServiceInstance.getUser( ctx, userID )
	result, err := userServiceInstance.getUserWithCache( ctx, userID )
	if err != nil {
		return nil, err
	}

	return result, nil
}


// ***************************************************************************************************
// ***************************************************************************************************
func getUserValidateInputData( userID uint64 ) error {
	if userID <= 0 {
		return errors.New("user ID is required")
	}

	return nil
}


// ***************************************************************************************************
// ***************************************************************************************************
func (userServiceInstance *userService) getUserWithCache(
	ctx context.Context,
	userID uint64,
) (*user_model.GetUserResponse, error) {
	cacheKey := getCacheKey( userID )

	// Проверяем в кэше
	values, err := userServiceInstance.cache.HGetAll(ctx, cacheKey)
	if err != nil {
		return nil, err
	}

	if len(values) != 0 {
		var userData user_model.GetUserResponseForRedis
		err = redis.ScanStruct(values, &userData)
		if err != nil {
			return nil, err
		}

		result := user_converter.GetUserWithCacheConvertFromRedis( userData )

		return result, nil
	}

	// Если в кэше нет, то получаем из БД
	result, err := userServiceInstance.getUser( ctx, userID )
	if err != nil {
		return nil, err
	}

	// Сохраняем в кэш
	cachePayload := user_converter.GetUserWithCacheConvertToRedis( result )
	err = userServiceInstance.cache.HashSet(ctx, cacheKey, cachePayload)
	if err != nil {
		return nil, err
	}

	return result, nil
}


// ***************************************************************************************************
// ***************************************************************************************************
func (userServiceInstance *userService) getUser(
	ctx context.Context,
	userID uint64,
) (*user_model.GetUserResponse, error) {
	userIdInt64 := int64(userID)

	userData, err := userServiceInstance.repository.GetUser(ctx, userIdInt64)
	if err != nil {
		return nil, err
	}

	result := user_converter.ToGetUserResponseFromRepository( &userData )

	return result, nil
}
