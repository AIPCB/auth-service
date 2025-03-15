package database

import (
	"context"
	"errors"

	"github.com/AIPCB/auth-service/src/sqlc"
)

type Client struct {
	queries *sqlc.Queries
}

func NewClient(opts ...Option) (*Client, error) {
	c := &Client{
		queries: nil,
	}

	for _, opt := range opts {
		opt(c)
	}

	if c.queries == nil {
		return nil, errors.New("queries are required")
	}

	return c, nil

}

func (c *Client) CreateUser(ctx context.Context, user sqlc.CreateUserParams) (sqlc.User, error) {
	return c.queries.CreateUser(ctx, user)
}
