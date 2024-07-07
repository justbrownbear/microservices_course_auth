package user_controller

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/justbrownbear/microservices_course_auth/pkg/user_v1"
)

func (s *controller) Update(ctx context.Context, req *user_v1.UpdateRequest) ( *emptypb.Empty, error ) {

	log.Printf("Update request fired: %v", req.String())

	payload := &emptypb.Empty{}

	return payload, nil
}
