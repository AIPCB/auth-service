package server

import (
	"github.com/AIPCB/auth-service/src/service"
	"github.com/AIPCB/auth-service/src/storage"
)

type Option func(*Server)

func WithAuthService(authService *service.Service) Option {
	return func(server *Server) {
		server.authService = authService
	}
}

func WithStorage(storage storage.Client) Option {
	return func(server *Server) {
		server.storage = storage
	}
}
