package service

import (
	"github.com/AIPCB/auth-service/src/storage"
)

type AuthService struct {
	storage storage.Client
}

func NewAuthService(options ...Option) *AuthService {
	service := &AuthService{}

	for _, option := range options {
		option(service)
	}

	return service
}
