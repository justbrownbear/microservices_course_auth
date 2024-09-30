package config

import (
	"errors"
	"os"
	"time"
)

const (
	redisHostEnvName                 = "REDIS_HOST"
	redisPortEnvName                 = "REDIS_PORT"
	redisConnectionTimeoutSecEnvName = "REDIS_CONNECTION_TIMEOUT_SEC"
	redisMaxIdleSecEnvName           = "REDIS_MAX_IDLE_SEC"
	redisIdleTimeoutSecEnvName       = "REDIS_IDLE_TIMEOUT_SEC"
)

type redisConfig struct {
	host                 string
	port                 uint16
	connectionTimeoutSec time.Duration
	maxIdleSec           int
	idleTimeoutSec       time.Duration
}

// RedisConfig интерфейс для получения конфигурации Redis
type RedisConfig interface {
	GetRedisHost() string
	GetRedisPort() uint16
	GetRedisConnectionTimeoutSec() time.Duration
	GetRedisMaxIdle() int
	GetRedisIdleTimeoutSec() time.Duration
}

// GetRedisConfig возвращает конфигурацию PostgreSQL
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

	connectionTimeoutSecInt, err := stringToInt(connectionTimeoutSec)
	if err != nil {
		return nil, err
	}
	connectionTimeoutSecDuration := time.Duration(connectionTimeoutSecInt) * time.Second

	maxIdleSec := os.Getenv(redisMaxIdleSecEnvName)
	if len(port) == 0 {
		return nil, errors.New(redisMaxIdleSecEnvName + " parameter not set")
	}

	maxIdleInt, err := stringToInt(maxIdleSec)
	if err != nil {
		return nil, err
	}

	idleTimeoutSec := os.Getenv(redisIdleTimeoutSecEnvName)
	if len(port) == 0 {
		return nil, errors.New(redisIdleTimeoutSecEnvName + " parameter not set")
	}

	idleTimeoutSecInt, err := stringToInt(idleTimeoutSec)
	if err != nil {
		return nil, err
	}
	idleTimeoutSecDuration := time.Duration(idleTimeoutSecInt) * time.Second

	result := &redisConfig{
		host:                 host,
		port:                 portUint16,
		connectionTimeoutSec: connectionTimeoutSecDuration,
		maxIdleSec:           maxIdleInt,
		idleTimeoutSec:       idleTimeoutSecDuration,
	}

	return result, nil
}

func (instance *redisConfig) GetRedisHost() string {
	return instance.host
}

func (instance *redisConfig) GetRedisPort() uint16 {
	return instance.port
}

func (instance *redisConfig) GetRedisConnectionTimeoutSec() time.Duration {
	return instance.connectionTimeoutSec
}

func (instance *redisConfig) GetRedisMaxIdle() int {
	return instance.maxIdleSec
}

func (instance *redisConfig) GetRedisIdleTimeoutSec() time.Duration {
	return instance.idleTimeoutSec
}
