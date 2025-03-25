package database

import (
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
