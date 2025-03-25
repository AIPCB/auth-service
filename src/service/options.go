package service

type Option func(*Service)

type Storage interface {
}

func WithStorage(s Storage) Option {
	return func(a *Service) {
		a.storage = s
	}
}
