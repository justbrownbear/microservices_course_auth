package config

import (
	"errors"
	"os"
)

const (
	redisHostEnvName					= "REDIS_HOST"
	redisPortEnvName					= "REDIS_PORT"
	redisConnectionTimeoutSecEnvName	= "REDIS_CONNECTION_TIMEOUT_SEC"
	redisMaxIdleSecEnvName				= "REDIS_MAX_IDLE_SEC"
	redisIdleTimeoutSecEnvName			= "REDIS_IDLE_TIMEOUT_SEC"
)


type redisConfig struct {
	host					string
	port					uint16
	connectionTimeoutSec	uint16
	maxIdleSec				uint16
	idleTimeoutSec			uint16
}


// PostgresqlConfig интерфейс для получения конфигурации PostgreSQL
type RedisConfig interface {
	GetRedisHost() string
	GetRedisPort() uint16
	GetRedisConnectionTimeoutSec() uint16
	GetRedisMaxIdle() uint16
	GetRedisIdleTimeoutSec() uint16
}

// GetPostgresqlConfig возвращает конфигурацию PostgreSQL
func GetRedisConfig() (RedisConfig, error) {
	host := os.Getenv(redisHostEnvName)
	if len(host) == 0 {
		return nil, errors.New(redisHostEnvName + " parameter not set")
	}

	port := os.Getenv(redisPortEnvName)
	if len(port) == 0 {
		return nil, errors.New(redisPortEnvName + " parameter not set")
	}

	portUint16, err := stringToUint16(port)
	if err != nil {
		return nil, err
	}

	connectionTimeoutSec := os.Getenv(redisConnectionTimeoutSecEnvName)
	if len(port) == 0 {
		return nil, errors.New(redisConnectionTimeoutSecEnvName + " parameter not set")
	}

	connectionTimeoutSecUint16, err := stringToUint16(connectionTimeoutSec)
	if err != nil {
		return nil, err
	}

	maxIdleSec := os.Getenv(redisMaxIdleSecEnvName)
	if len(port) == 0 {
		return nil, errors.New(redisMaxIdleSecEnvName + " parameter not set")
	}

	maxIdleUint16, err := stringToUint16(maxIdleSec)
	if err != nil {
		return nil, err
	}

	idleTimeoutSec := os.Getenv(redisIdleTimeoutSecEnvName)
	if len(port) == 0 {
		return nil, errors.New(redisIdleTimeoutSecEnvName + " parameter not set")
	}

	idleTimeoutSecUint16, err := stringToUint16(idleTimeoutSec)
	if err != nil {
		return nil, err
	}


	result := &redisConfig{
		host:					host,
		port:					portUint16,
		connectionTimeoutSec:	connectionTimeoutSecUint16,
		maxIdleSec:				maxIdleUint16,
		idleTimeoutSec:			idleTimeoutSecUint16,
	}

	return result, nil
}

func (c *redisConfig) GetRedisHost() string {
	return c.host
}

func (c *redisConfig) GetRedisPort() uint16 {
	return c.port
}

func (c *redisConfig) GetRedisConnectionTimeoutSec() uint16 {
	return c.connectionTimeoutSec
}

func (c *redisConfig) GetRedisMaxIdle() uint16 {
	return c.maxIdleSec
}

func (c *redisConfig) GetRedisIdleTimeoutSec() uint16 {
	return c.idleTimeoutSec
}
