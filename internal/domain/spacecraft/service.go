package spacecraft

import (
	"context"

	"github.com/google/uuid"
)

type (
	repo interface {
		List(context.Context) ([]Spacecraft, error)
		Get(context.Context, uuid.UUID) (Spacecraft, error)
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

func (s *Service) Get(ctx context.Context, id uuid.UUID) (Spacecraft, error) {
	spacecraft, err := s.repo.Get(ctx, id)
	if err != nil {
		return Spacecraft{}, err
	}

	return spacecraft, nil
}
