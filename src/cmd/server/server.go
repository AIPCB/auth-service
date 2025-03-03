package cmd

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/AIPCB/auth-service/src/models"
	"github.com/AIPCB/auth-service/src/repo"
	"github.com/AIPCB/auth-service/src/server"
	"github.com/AIPCB/auth-service/src/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Execute() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// TODO: abstract, use env variables, use secure connection
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to open database: %+v", err)
	}

	db.AutoMigrate(&models.User{})

	authService := service.NewAuthService(service.WithRepo(repo.NewPostgresRepo(db)))

	s := server.NewServer(
		server.WithAuthService(authService),
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
