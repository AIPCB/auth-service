package service

import (
	"context"

	"github.com/AIPCB/auth-service/src/sqlc"
)

type Option func(*Service)

type Storage interface {
	CreateUser(ctx context.Context, user sqlc.CreateUserParams) (sqlc.User, error)
}

func WithStorage(s Storage) Option {
	return func(a *Service) {
		a.storage = s
	}
}
