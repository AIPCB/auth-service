package server

import (
	"github.com/AIPCB/auth-service/src/service"
)

type Option func(*Server)

func WithAuthService(authService *service.AuthService) Option {
	return func(server *Server) {
		server.authService = authService
	}
}
