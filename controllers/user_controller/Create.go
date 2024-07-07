package user_controller

import (
	"context"
	"log"
	"math/rand"

	"github.com/justbrownbear/microservices_course_auth/pkg/user_v1"
)



func (s *controller) Create(ctx context.Context, req *user_v1.CreateRequest) (*user_v1.CreateResponse, error) {

	log.Printf("Create request fired: %v", req.String())

	payload := &user_v1.CreateResponse{
		Id: rand.Int63n( 100500 ),
	}

	return payload, nil
}
