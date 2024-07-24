package app

import (
	"log"
	"net"
	"strconv"

	// "github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/justbrownbear/microservices_course_auth/controllers/user_controller"
)

var grpcServer *grpc.Server

// InitApp initializes the gRPC server and registers the user controller.
// It also enables server reflection for easier debugging and service discovery.
func InitApp() {
	grpcServer = grpc.NewServer()
	reflection.Register(grpcServer)

	user_controller.InitUserController(grpcServer)
}

// StartApp initializes a gRPC server listener on the specified protocol and port.
// It logs an error message if the listener fails to initialize.
// On successful initialization, it logs a message indicating the gRPC server is starting on the specified address.
func StartApp(grpcProtocol string, grpcPort uint16) error {
	listenAddress := ":" + strconv.Itoa(int(grpcPort))

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
