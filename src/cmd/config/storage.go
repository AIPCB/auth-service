package config

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/AIPCB/auth-service/src/storage"
)

func NewStorageClient(ctx context.Context) (*storage.Client, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSLMODE"))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)

	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	database, err := NewDatabaseStore(ctx, db)
	if err != nil {
		db.Close()
		return nil, err
	}

	return storage.NewClient(
		storage.WithDatabase(database),
	)
}
