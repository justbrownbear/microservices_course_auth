package main

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/justbrownbear/microservices_course_auth/app"
)


const GRPC_PORT = 9099;



func main() {
	app.InitApp()
	err := app.StartApp( GRPC_PORT );

	if err != nil {
		fmt.Println( color.RedString( "Failed to start app: %v", err ) )
	}
}
