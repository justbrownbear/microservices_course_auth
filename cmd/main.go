package main

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/justbrownbear/microservices_course_auth/app"

)


const gRpcPort = 9099;



func main() {
	app.InitApp()
	err := app.StartApp( gRpcPort );

	if err != nil {
		fmt.Println( color.RedString( "Failed to start app: %v", err ) )
	}
}
