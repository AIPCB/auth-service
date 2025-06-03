package person

import (
	"context"
	"errors"
	"fmt"
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

func (s *PersonService) GetPerson(ctx context.Context, id string) (*models.Person, error) {
	resp := &models.Person{}
	url := fmt.Sprintf("person/%s", id)
	err := s.client.DoRequest(ctx, "GET", url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *PersonService) GetPersonByEmail(ctx context.Context, email string) (*models.Person, error) {
	resp := &models.Person{}
	url := fmt.Sprintf("person/by-email/%s", url.QueryEscape(email))
	err := s.client.DoRequest(ctx, "GET", url, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
