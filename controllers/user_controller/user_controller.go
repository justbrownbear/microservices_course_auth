package user_controller

import (
	"github.com/justbrownbear/microservices_course_auth/pkg/user_v1"
	"google.golang.org/grpc"
)


type controller struct {
	user_v1.UnimplementedUserV1Server
}


func InitUserController( grpcServer *grpc.Server ) {

	user_v1.RegisterUserV1Server( grpcServer, &controller{} )
}
