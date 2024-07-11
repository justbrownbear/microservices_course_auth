package user_controller

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"log"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/justbrownbear/microservices_course_auth/pkg/user_v1"
)

func (s *controller) Get(_ context.Context, req *user_v1.GetRequest) (*user_v1.GetResponse, error) {

	log.Printf("Get request fired: %v", req.String())

	createdUpdatedDate := gofakeit.Date()

	result := &user_v1.GetResponse{
		Id:        req.GetId(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      generateRandomRole(),
		CreatedAt: timestamppb.New(createdUpdatedDate),
		UpdatedAt: timestamppb.New(createdUpdatedDate),
	}

	return result, nil
}

func generateRandomRole() user_v1.Role {
	var result int32

	err := binary.Read(rand.Reader, binary.BigEndian, &result)

	if err != nil {
		log.Printf("Failed to generate random id: %v", err)
		return 0
	}

	if result < 0 {
		result = -result
	}

	result = result % 2
	role := user_v1.Role(result)

	return role
}
