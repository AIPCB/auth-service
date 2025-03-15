package storage

import "github.com/AIPCB/auth-service/src/storage/database"

type Option func(*Client)

func WithDatabase(db *database.Client) Option {
	return func(c *Client) {
		c.db = db
	}
}
