package service_provider

import (
	"github.com/jackc/pgx/v5"
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
func NewWithTransaction(dbTransaction *pgx.Tx) ServiceProvider {
	return &serviceProvider{
		dbTransaction: dbTransaction,
	}
}


// ***************************************************************************************************
// ***************************************************************************************************
func (serviceProviderInstance *serviceProvider) getUserRepository() user_repository.UserRepository {
	if serviceProviderInstance.userRepository == nil {
		serviceProviderInstance.userRepository = user_repository.New(serviceProviderInstance.dbConnection)

		if serviceProviderInstance.dbTransaction != nil {
			serviceProviderInstance.userRepository =
				serviceProviderInstance.userRepository.WithTx(*serviceProviderInstance.dbTransaction)
		}
	}

	return serviceProviderInstance.userRepository
}


// ***************************************************************************************************
// ***************************************************************************************************
func (serviceProviderInstance *serviceProvider) GetUserService() user_service.UserService {
	if serviceProviderInstance.userService == nil {
		serviceProviderInstance.userService =
			user_service.New(serviceProviderInstance.getUserRepository())
	}

	return serviceProviderInstance.userService
}
