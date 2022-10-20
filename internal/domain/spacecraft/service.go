package spacecraft

import (
	"context"
)

type (
	repo interface {
		List(context.Context) ([]Spacecraft, error)
	}

	Service struct {
		repo repo
	}
)

func NewService(repo repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) List(ctx context.Context) ([]Spacecraft, error) {
	spacecrafts, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	return spacecrafts, nil
}
