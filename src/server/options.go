package server

import (
	"time"

	"github.com/AIPCB/auth-service/src/service"
	"github.com/AIPCB/auth-service/src/storage"
)

type Option func(*Server)

type Service interface {
}

func WithAuthService(authService *service.Service) Option {
	return func(server *Server) {
		server.authService = authService
	}
}

func WithPersonService(personService service.PersonService) Option {
	return func(server *Server) {
		server.personService = personService
	}
}

func WithStorage(storage storage.Client) Option {
	return func(server *Server) {
		server.storage = storage
	}
}

func WithJWTSecret(secret string) Option {
	return func(server *Server) {
		server.jwtSecret = []byte(secret)
	}
}

func WithJWTExpiryTime(expiryTime time.Duration) Option {
	return func(server *Server) {
		server.jwtExpiryTime = expiryTime
	}
}
