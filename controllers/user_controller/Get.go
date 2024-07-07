package user_controller

import (
	"context"
	"log"
	"math/rand"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/justbrownbear/microservices_course_auth/pkg/user_v1"
)



func (s *controller) Get(ctx context.Context, req *user_v1.GetRequest) (*user_v1.GetResponse, error) {

	log.Printf("Get request fired: %v", req.String())

	createdUpdatedDate := gofakeit.Date()

	result := &user_v1.GetResponse{
		Id:			req.GetId(),
		Name:		gofakeit.Name(),
		Email:		gofakeit.Email(),
		Role:		user_v1.Role( rand.Intn( 2 ) ),
		CreatedAt:	timestamppb.New( createdUpdatedDate ),
		UpdatedAt:	timestamppb.New( createdUpdatedDate ),
	}

	return result, nil
}
