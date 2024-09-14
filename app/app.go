package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/fatih/color"
	"github.com/gomodule/redigo/redis"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/justbrownbear/microservices_course_auth/controllers/user_controller"
	grpc_api "github.com/justbrownbear/microservices_course_auth/internal/api/grpc"
	"github.com/justbrownbear/microservices_course_auth/internal/config"
	"github.com/justbrownbear/microservices_course_auth/internal/transaction_manager"
)

var dbPool *pgxpool.Pool
var redisPool *redis.Pool
var grpcServer *grpc.Server

var grpcConfig config.GRPCConfig

// InitApp initializes the gRPC server and registers the user controller.
// It also enables server reflection for easier debugging and service discovery.
func InitApp(
	ctx context.Context,
	postgresqlConfig config.PostgresqlConfig,
	redisConfig config.RedisConfig,
	grpcConfigInstance config.GRPCConfig,
) error {
	grpcConfig = grpcConfigInstance
	grpcServer = grpc.NewServer()
	reflection.Register(grpcServer)

	var err error

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
	grpcAPI := grpc_api.InitGrpcAPI(transactionManager)

	user_controller.InitUserController(grpcServer, grpcAPI)

	return nil
}

// StartApp initializes a gRPC server listener on the specified protocol and port.
// It logs an error message if the listener fails to initialize.
// On successful initialization, it logs a message indicating the gRPC server is starting on the specified address.
func StartApp() error {
	grpcProtocol := grpcConfig.GetGrpcProtocol()
	grpcHost := grpcConfig.GetGrpcHost()
	grpcPort := grpcConfig.GetGrpcPort()

	listenAddress := net.JoinHostPort(grpcHost, strconv.Itoa(int(grpcPort)))
	listener, err := net.Listen(grpcProtocol, listenAddress)
	if err != nil {
		log.Printf(color.RedString("Failed to initialize listener: %v"), err)
		return err
	}

	log.Printf(color.GreenString("Starting gRPC server on %s"), listenAddress)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Printf(color.RedString("Failed to start gRPC server: %v"), err)
		return err
	}

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

	log.Println(color.GreenString("Application stopped successfully. Bye."))
}
