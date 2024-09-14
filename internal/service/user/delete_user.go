package user_service

import (
	"context"
	"errors"
)

// ***************************************************************************************************
// ***************************************************************************************************
func (instance *userService) DeleteUser(ctx context.Context, userID uint64) error {
	err := deleteUserValidateInputData(userID)
	if err != nil {
		return err
	}

	err = instance.deleteUserWithCache(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}


// ***************************************************************************************************
// ***************************************************************************************************
func deleteUserValidateInputData( userID uint64 ) error {
	if userID <= 0 {
		return errors.New("user ID is required")
	}

	return nil
}

// ***************************************************************************************************
// ***************************************************************************************************
func (instance *userService) deleteUserWithCache(ctx context.Context, userID uint64) error {
	err := instance.deleteUser(ctx, userID)
	if err != nil {
		return err
	}

	// Очищаем кэш пользователя
	cacheKey := getCacheKey( userID )
	err = instance.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}

	return nil
}

// ***************************************************************************************************
// ***************************************************************************************************
func (instance *userService) deleteUser(ctx context.Context, userID uint64) error {
	userIdInt64 := int64(userID)

	err := instance.repository.DeleteUser(ctx, userIdInt64)
	if err != nil {
		return err
	}

	return nil
}
