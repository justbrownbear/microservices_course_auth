package config

import (
	"errors"
	"os"
	"strings"
)

const (
	kafkaBrokersEnvName = "KAFKA_BROKERS"
	kafkaGroupIDEnvName = "KAFKA_GROUP_ID"
)

type kafkaConfig struct {
	brokers []string
	groupID string
}

// KafkaConfig интерфейс для получения конфигурации Kafka
type KafkaConfig interface {
	GetBrokers() []string
	GetGroupID() string
}

// GetKafkaConfig возвращает конфигурацию Kafka
func GetKafkaConfig() (KafkaConfig, error) {
	brokersString := os.Getenv(kafkaBrokersEnvName)
	if len(brokersString) == 0 {
		return nil, errors.New(kafkaBrokersEnvName + " parameter not set")
	}

	groupID := os.Getenv(kafkaGroupIDEnvName)
	if len(groupID) == 0 {
		return nil, errors.New(kafkaGroupIDEnvName + " parameter not set")
	}

	brokers := strings.Split(brokersString, ",")
	// Подчистим пробелы
	for i, broker := range brokers {
		brokers[i] = strings.TrimSpace(broker)
	}

	result := &kafkaConfig{
		brokers,
		groupID,
	}

	return result, nil
}

func (instance *kafkaConfig) GetBrokers() []string {
	return instance.brokers
}

func (instance *kafkaConfig) GetGroupID() string {
	return instance.groupID
}
