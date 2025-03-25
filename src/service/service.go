package service

import (
	"errors"
)

type Service struct {
	storage Storage
}

func NewService(options ...Option) (*Service, error) {
	s := &Service{}

	for _, option := range options {
		option(s)
	}

	if s.storage == nil {
		return nil, errors.New("service: missing storage")
	}

	return s, nil
}
