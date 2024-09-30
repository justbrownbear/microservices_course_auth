package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"

	"github.com/IBM/sarama"
	"github.com/fatih/color"
	"github.com/gomodule/redigo/redis"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	user_controller "github.com/justbrownbear/microservices_course_auth/controllers/user_grpc_controller"
	"github.com/justbrownbear/microservices_course_auth/controllers/user_kafka_controller"
	grpc_api "github.com/justbrownbear/microservices_course_auth/internal/api/grpc"
	kafka_api "github.com/justbrownbear/microservices_course_auth/internal/api/kafka"
	"github.com/justbrownbear/microservices_course_auth/internal/client/kafka"
	kafka_consumer "github.com/justbrownbear/microservices_course_auth/internal/client/kafka/consumer"
	"github.com/justbrownbear/microservices_course_auth/internal/config"
	"github.com/justbrownbear/microservices_course_auth/internal/transaction_manager"
)

var dbPool *pgxpool.Pool
var redisPool *redis.Pool
var grpcServer *grpc.Server

var grpcConfig config.GRPCConfig
var kafkaConfig config.KafkaConfig

var kafkaAPI kafka_api.KafkaAPI
var kafkaController user_kafka_controller.UserKafkaController
var kafkaConsumer kafka.Consumer
var kafkaProducer sarama.SyncProducer

// InitApp initializes the gRPC server and registers the user controller.
// It also enables server reflection for easier debugging and service discovery.
func InitApp(
	ctx context.Context,
	postgresqlConfig config.PostgresqlConfig,
	redisConfig config.RedisConfig,
	grpcConfigInstance config.GRPCConfig,
	kafkaConfigInstance config.KafkaConfig,
) error {
	var err error

	grpcConfig = grpcConfigInstance
	grpcServer = initGrpcServer()

	kafkaConfig = kafkaConfigInstance
	kafkaConsumer = initKafkaConsumer()

	kafkaProducer, err = initKafkaProducer()
	if err != nil {
		log.Printf(color.RedString("Failed to initialize Kafka producer: %v"), err)
		return err
	}

	dbDSN :=
		fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
			postgresqlConfig.GetPostgresHost(),
			postgresqlConfig.GetPostgresPort(),
			postgresqlConfig.GetPostgresDb(),
			postgresqlConfig.GetPostgresUser(),
			postgresqlConfig.GetPostgresPassword())
	dbPool, err = pgxpool.New(ctx, dbDSN)
	if err != nil {
		log.Printf(color.RedString("Unable to connect to database: %v\n"), err)
		return err
	}

	redisPool = &redis.Pool{
		MaxIdle:     redisConfig.GetRedisMaxIdle(),
		IdleTimeout: redisConfig.GetRedisIdleTimeoutSec(),
		DialContext: func(ctx context.Context) (redis.Conn, error) {
			redisAddress := net.JoinHostPort(redisConfig.GetRedisHost(), strconv.Itoa(int(redisConfig.GetRedisPort())))
			return redis.DialContext(ctx, "tcp", redisAddress)
		},
	}

	transactionManager := transaction_manager.InitTransactionManager(dbPool, redisPool, &redisConfig)
	grpcAPI := grpc_api.InitGrpcAPI(transactionManager, kafkaProducer)
	kafkaAPI = kafka_api.InitKafkaAPI(transactionManager)

	user_controller.InitUserController(grpcServer, grpcAPI)
	kafkaController = user_kafka_controller.InitUserKafkaController(kafkaAPI)

	return nil
}

// StartApp initializes a gRPC server listener on the specified protocol and port.
// It logs an error message if the listener fails to initialize.
// On successful initialization, it logs a message indicating the gRPC server is starting on the specified address.
func StartApp(ctx context.Context) error {
	waitGroup := sync.WaitGroup{}

	waitGroup.Add(1)
	go startGrpcServer(&waitGroup)
	waitGroup.Add(1)
	go startKafkaConsumer(ctx, &waitGroup)

	waitGroup.Wait()

	return nil
}

// StopApp - Остановка приложения
func StopApp() {
	log.Println(color.YellowString("Stopping the application right way..."))

	grpcServer.Stop()
	dbPool.Close()
	err := redisPool.Close()
	if err != nil {
		log.Printf(color.RedString("Failed to close redis pool: %v"), err)
	}

	err = kafkaConsumer.Close()
	if err != nil {
		log.Printf(color.RedString("Failed to close Kafka consumer: %v"), err)
	}

	log.Println(color.GreenString("Application stopped successfully. Bye."))
}

// ***************************************************************************************************
// ***************************************************************************************************
func initGrpcServer() *grpc.Server {
	grpcServerInstance := grpc.NewServer()
	reflection.Register(grpcServerInstance)

	return grpcServerInstance
}

// ***************************************************************************************************
// ***************************************************************************************************
func startGrpcServer(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	grpcProtocol := grpcConfig.GetGrpcProtocol()
	listenAddress := getGrpcAddress()
	listener, err := net.Listen(grpcProtocol, listenAddress)
	if err != nil {
		log.Printf(color.RedString("Failed to initialize listener: %v"), err)
		return
	}

	log.Printf(color.GreenString("Starting gRPC server on %s"), listenAddress)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Printf(color.RedString("Failed to start gRPC server: %v"), err)
		return
	}
}

// ***************************************************************************************************
// ***************************************************************************************************
func getGrpcAddress() string {
	grpcHost := grpcConfig.GetGrpcHost()
	grpcPort := grpcConfig.GetGrpcPort()

	listenAddress := net.JoinHostPort(grpcHost, strconv.Itoa(int(grpcPort)))

	return listenAddress
}

// ***************************************************************************************************
// ***************************************************************************************************
func initKafkaConsumer() kafka.Consumer {
	consumerGroup, err := sarama.NewConsumerGroup(
		kafkaConfig.GetBrokers(),
		kafkaConfig.GetGroupID(),
		nil,
	)
	if err != nil {
		log.Fatalf("failed to create consumer group: %v", err)
	}

	consumerGroupHandler := kafka_consumer.NewGroupHandler()
	consumer := kafka_consumer.NewConsumer(
		consumerGroup,
		consumerGroupHandler,
	)

	return consumer
}

// ***************************************************************************************************
// ***************************************************************************************************
func startKafkaConsumer(ctx context.Context, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	err := kafkaConsumer.Consume(ctx, kafka.KafkaTopicUser, kafkaController.OnMessage)
	if err != nil {
		log.Printf(color.RedString("Failed to consume messages: %v"), err)
		return
	}
}

// ***************************************************************************************************
// ***************************************************************************************************
func initKafkaProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(kafkaConfig.GetBrokers(), config)
	if err != nil {
		return nil, err
	}

	return producer, nil
}
