package user_service

import "strconv"

var cacheKeyPrefix = "user_service:GetUser:"

func getCacheKey(userID uint64) string {
	return cacheKeyPrefix + strconv.Itoa(int(userID))
}
