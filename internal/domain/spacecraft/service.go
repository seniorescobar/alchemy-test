package spacecraft

import (
	"context"

	"github.com/google/uuid"
)

type (
	repo interface {
		List(context.Context) ([]Spacecraft, error)
		Get(context.Context, uuid.UUID) (Spacecraft, error)
		Create(context.Context, Spacecraft) error
		Update(context.Context, Spacecraft) error
		Delete(context.Context, uuid.UUID) error
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

func (s *Service) Create(ctx context.Context, spacecraft Spacecraft) (Spacecraft, error) {
	spacecraft.ID = uuid.New()

	if err := s.repo.Create(ctx, spacecraft); err != nil {
		return Spacecraft{}, err
	}

	return spacecraft, nil
}

func (s *Service) Update(ctx context.Context, spacecraft Spacecraft) (Spacecraft, error) {
	if err := s.repo.Update(ctx, spacecraft); err != nil {
		return Spacecraft{}, err
	}

	return spacecraft, nil
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
