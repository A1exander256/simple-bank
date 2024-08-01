package user

import (
	"context"

	"github.com/google/uuid"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateUser(_ context.Context, _ *CreateUserParams) (uuid.UUID, error) {
	return uuid.New(), nil
}
