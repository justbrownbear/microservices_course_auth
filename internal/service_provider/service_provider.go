package service_provider

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jackc/pgx/v5"
	"github.com/justbrownbear/microservices_course_auth/internal/client/cache"
	redis_cache "github.com/justbrownbear/microservices_course_auth/internal/client/cache/redis"
	"github.com/justbrownbear/microservices_course_auth/internal/config"
	user_repository "github.com/justbrownbear/microservices_course_auth/internal/repository/user"
	user_service "github.com/justbrownbear/microservices_course_auth/internal/service/user"
)

// ServiceProvider - Интерфейс сервис-провайдера
type ServiceProvider interface {
	GetUserService() user_service.UserService
}


type serviceProvider struct {
	dbConnection	*pgx.Conn
	dbTransaction	*pgx.Tx

	redisPool		*redis.Pool
	redisConfig		*config.RedisConfig

	redisClient		cache.RedisClient

	userRepository	user_repository.UserRepository
	userService		user_service.UserService
}


// ***************************************************************************************************
// ***************************************************************************************************
// NewWithConnection создает новый экземпляр сервис-провайдера с соединением
func NewWithConnection(dbConnection *pgx.Conn) ServiceProvider {
	return &serviceProvider{
		dbConnection: dbConnection,
	}
}


// ***************************************************************************************************
// ***************************************************************************************************
// NewWithTransaction создает новый экземпляр сервис-провайдера с транзакцией
func NewWithTransaction(dbTransaction *pgx.Tx, redisPool *redis.Pool, redisConfig *config.RedisConfig) ServiceProvider {
	return &serviceProvider{
		dbTransaction: dbTransaction,
		redisPool: redisPool,
		redisConfig: redisConfig,
	}
}


// ***************************************************************************************************
// ***************************************************************************************************
func (instance *serviceProvider) getCacheClient() cache.RedisClient {
	if instance.redisClient == nil {
		instance.redisClient = redis_cache.NewClient( instance.redisPool, *instance.redisConfig )
	}

	return instance.redisClient
}


// ***************************************************************************************************
// ***************************************************************************************************
func (instance *serviceProvider) getUserRepository() user_repository.UserRepository {
	if instance.userRepository == nil {
		instance.userRepository = user_repository.New(instance.dbConnection)

		if instance.dbTransaction != nil {
			instance.userRepository =
				instance.userRepository.WithTx(*instance.dbTransaction)
		}
	}

	return instance.userRepository
}


// ***************************************************************************************************
// ***************************************************************************************************
func (instance *serviceProvider) GetUserService() user_service.UserService {
	if instance.userService == nil {
		instance.userService =
			user_service.New(instance.getUserRepository(), instance.getCacheClient())
	}

	return instance.userService
}
