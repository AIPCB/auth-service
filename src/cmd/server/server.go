package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AIPCB/auth-service/src/cmd/config"
	"github.com/AIPCB/auth-service/src/server"
	"github.com/AIPCB/auth-service/src/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// TODO: Break up into smaller pieces
func Execute() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	storageClient, err := config.NewStorageClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create storage client: %+v", err)
	}

	authService, err := service.NewService(
		service.WithStorage(storageClient),
	)

	if err != nil {
		log.Fatalf("Failed to create auth service: %+v", err)
		return
	}

	s := server.NewServer(
		server.WithAuthService(authService),
		server.WithJWTExpiryTime(time.Hour*24),
		server.WithJWTSecret(os.Getenv("JWT_SECRET")),
	)

	go func() {
		log.Println("Starting server on port 8080...")

		if err := s.ListenAndServe(":8080"); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down server...")
	if err := s.Shutdown(context.Background()); err != nil {
		log.Fatalf("Server shutdown failed: %+v", err)
	}
	log.Println("Server gracefully stopped")
}
