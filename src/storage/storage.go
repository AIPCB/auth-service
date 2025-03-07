package storage

import (
	"context"
	"database/sql"
)

type Client interface {
	CreateUser(ctx context.Context, email, username, password string) (User, error)
}

type Storage struct {
	queries *Queries
}

func NewStorageClient(db *sql.DB) *Storage {
	return &Storage{queries: New(db)}
}

func (s *Storage) CreateUser(ctx context.Context, email, username, password string) (User, error) {
	return s.queries.CreateUser(
		ctx,
		CreateUserParams{
			Email:    email,
			Username: username,
			Password: password,
		},
	)
}
