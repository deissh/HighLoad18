package filters

import "github.com/deissh/HighLoad18/core/account"

type Repository interface {
	Fetch(filters []Filter, limit int) ([]account.Account, error)
}

type Service interface {
	Fetch(filters []Filter, limit int) ([]account.Account, error)
}

type service struct {
	repo Repository
}

func New(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) Fetch(filters []Filter, limit int) ([]account.Account, error) {
	return s.repo.Fetch(filters, limit)
}
