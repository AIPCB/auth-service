package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AIPCB/auth-service/src/server"
	"github.com/AIPCB/auth-service/src/service"
	"github.com/AIPCB/auth-service/src/storage"
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

	// TODO: abstract, use secure connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSLMODE"))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %+v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to open database: %+v", err)
	}

	authService := service.NewAuthService()

	s := server.NewServer(
		server.WithAuthService(authService),
		server.WithStorage(
			storage.NewStorageClient(db),
		),
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
