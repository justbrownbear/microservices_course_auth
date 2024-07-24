package main

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/justbrownbear/microservices_course_auth/app"
)

// Used gRPC protocol
const gRPCProtocol = "tcp"

// Used gRPC port
const gRPCPort = 9099

func main() {
	app.InitApp()
	err := app.StartApp(gRPCProtocol, gRPCPort)

	if err != nil {
		fmt.Println(color.RedString("Failed to start app: %v", err))
	}
}
