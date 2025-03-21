package service

import (
	"context"
	"errors"

	"github.com/AIPCB/auth-service/src/sqlc"
)

type Service struct {
	storage Storage
}

func NewService(options ...Option) (*Service, error) {
	s := &Service{}

	for _, option := range options {
		option(s)
	}

	if s.storage == nil {
		return nil, errors.New("service: missing storage")
	}

	return s, nil
}

// TODO: move service related logic here

func (s *Service) RegisterUser(ctx context.Context, user sqlc.CreateUserParams) (sqlc.User, error) {
	return s.storage.CreateUser(ctx, user)
}
