package storage

import (
	"fmt"

	"github.com/AIPCB/auth-service/src/storage/database"
)

type Client struct {
	db *database.Client
}

func NewClient(opts ...Option) (*Client, error) {
	c := &Client{}

	for _, opt := range opts {
		opt(c)
	}

	if c.db == nil {
		return nil, fmt.Errorf("storage: missing queries in options")
	}

	return c, nil
}
