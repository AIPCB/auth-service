package person

import (
	"context"
	"errors"
	"net/url"

	"github.com/AIPCB/auth-service/src/models"
	httpclient "github.com/AIPCB/shared/src/http"
)

type PersonService struct {
	url    *url.URL
	client *httpclient.Client
}

func NewPersonService(opts ...Option) (*PersonService, error) {
	s := &PersonService{}
	for _, opt := range opts {
		opt(s)
	}

	if s.url == nil {
		return nil, errors.New("service: missing PersonService URL")
	}

	s.client = httpclient.New(s.url.String())

	return s, nil
}

func (s *PersonService) CreatePerson(ctx context.Context, req models.RegisterRequest) error {
	resp := &models.RegisterResponse{}
	err := s.client.DoRequest(ctx, "POST", "create-person", req, resp)
	if err != nil {
		return err
	}

	return nil
}
