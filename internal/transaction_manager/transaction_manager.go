package transaction_manager

import (
	"context"

	"github.com/gomodule/redigo/redis"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"

	"github.com/justbrownbear/microservices_course_auth/internal/config"
	"github.com/justbrownbear/microservices_course_auth/internal/service_provider"
)

// Handler - функция, которая выполняется в транзакции или коннекшне,
// это будет лежать в оригинале в другом package
type Handler func(ctx context.Context, serviceProvider service_provider.ServiceProvider) error

// TxManager defines an interface for managing database transactions.
// It provides methods to execute handlers within the context of a connection,
// either with or without a transaction.
type TxManager interface {
	// Выполнение handler в контексте подключения без транзакции
	// WithConnection(ctx context.Context, handler Handler) error
	// Выполнение handler в контексте подключения с транзакцией
	WithTransaction(ctx context.Context, handler Handler) error
}

// В перспективе тут может быть не только pg, но и другие подключения.
// Например, redis
type resources struct {
	dbPool      *pgxpool.Pool
	redisPool   *redis.Pool
	redisConfig *config.RedisConfig
}

// InitTransactionManager initializes and returns a new transaction manager instance.
// It takes a PostgreSQL connection pool, a Redis connection pool, and a Redis configuration as parameters.
//
// Parameters:
//   - dbPool: A pointer to a pgxpool.Pool representing the PostgreSQL connection pool.
//   - redisPool: A pointer to a redis.Pool representing the Redis connection pool.
//   - redisConfig: A pointer to a config.RedisConfig containing the Redis configuration.
//
// Returns:
//   - TxManager: An instance of the transaction manager.
func InitTransactionManager(dbPool *pgxpool.Pool, redisPool *redis.Pool, redisConfig *config.RedisConfig) TxManager {
	return &resources{
		dbPool:      dbPool,
		redisPool:   redisPool,
		redisConfig: redisConfig,
	}
}

// Разумеется, это будет лежать в оригинале в другом package
func (instance *resources) WithTransaction(
	ctx context.Context,
	handler Handler,
) error {
	// Инициализируем соединение
	transaction, err := instance.dbPool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return err
	}

	// Инициализируем сервис-провайдер транзакцией
	serviceProvider := service_provider.NewWithTransaction(&transaction, instance.redisPool, instance.redisConfig)
	// А можем и с коннекшеном
	// serviceProvider := getServiceProviderWithConnection(&connection)

	// TODO: В рамках менеджера транзакций мы так же эмулируем транзакции к Redis
	// redisClient := serviceProvider.getRedisClient()

	// Настраиваем функцию отсрочки для отката или коммита транзакции.
	defer func() {
		// Восстанавливаемся после паники
		recoverResult := recover()
		if recoverResult != nil {
			err = errors.Errorf("panic recovered: %v", recoverResult)
		}

		// Откатываем транзакцию, если произошла ошибка
		if err != nil {
			errRollback := transaction.Rollback(ctx)
			if errRollback != nil {
				err = errors.Wrapf(err, "pg errRollback: %v", errRollback)
			}

			// TODO: В рамках менеджера транзакций мы так же эмулируем транзакции к Redis
			// errRollback = redisClient.Close()
			// if errRollback != nil {
			// 	err = errors.Wrapf(err, "redis errRollback: %v", errRollback)
			// }

			return
		}
	}()

	// Выполняем бизнес-логику
	err = handler(ctx, serviceProvider)
	if err != nil {
		err = errors.Wrap(err, "failed executing code inside transaction")
	}

	// если ошибок не было, коммитим транзакцию
	if err == nil {
		err = transaction.Commit(ctx)
		if err != nil {
			err = errors.Wrap(err, "pg tx commit failed")
		}

		// TODO: В рамках менеджера транзакций мы так же эмулируем транзакции к Redis
		// err = redisClient.Commit(ctx)
		// if err != nil {
		// 	err = errors.Wrap(err, "redis tx commit failed")
		// }
	}

	return err
}
