package person

import (
	"net/url"
)

type Option func(*PersonService)

func WithPersonServiceURL(url *url.URL) Option {
	return func(s *PersonService) {
		s.url = url
	}
}
