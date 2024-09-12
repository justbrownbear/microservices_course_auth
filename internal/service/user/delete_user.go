package user_service

import (
	"context"
	"errors"
)

// ***************************************************************************************************
// ***************************************************************************************************
func (userServiceInstance *userService) DeleteUser(ctx context.Context, userID uint64) error {
	err := deleteUserValidateInputData(userID)
	if err != nil {
		return err
	}

	userIdInt64 := int64(userID)

	err = userServiceInstance.repository.DeleteUser(ctx, userIdInt64)
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
