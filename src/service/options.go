package service

import (
	"context"

	"github.com/AIPCB/auth-service/src/models"
)

type Option func(*Service)

type Storage interface {
}

type PersonService interface {
	CreatePerson(ctx context.Context, req models.RegisterRequest) error
	GetPersonByEmail(ctx context.Context, email string) (*models.Person, error)
}

func WithStorage(s Storage) Option {
	return func(a *Service) {
		a.storage = s
	}
}
