package spacecraft

import (
	"context"
	"net/url"
)

type (
	repo interface {
		List(context.Context) ([]Spacecraft, error)
		Get(context.Context, int) (Spacecraft, error)
		Create(context.Context, Spacecraft) error
		Update(context.Context, Spacecraft) error
		Delete(context.Context, int) error
	}

	Service struct {
		repo repo
	}
)

var (
	ErrInvalidImage  = &ValidationErr{"invalid image"}
	ErrInvalidStatus = &ValidationErr{"invalid status"}
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

func (s *Service) Get(ctx context.Context, id int) (Spacecraft, error) {
	spacecraft, err := s.repo.Get(ctx, id)
	if err != nil {
		return Spacecraft{}, err
	}

	return spacecraft, nil
}

func (s *Service) Create(ctx context.Context, spacecraft Spacecraft) error {
	if err := validateSpacecraft(spacecraft); err != nil {
		return err
	}

	return s.repo.Create(ctx, spacecraft)
}

func (s *Service) Update(ctx context.Context, spacecraft Spacecraft) error {
	if err := validateSpacecraft(spacecraft); err != nil {
		return err
	}

	return s.repo.Update(ctx, spacecraft)
}

func (s *Service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func validateSpacecraft(spacecraft Spacecraft) error {
	if err := validateImage(spacecraft.Image); err != nil {
		return err
	}

	if err := validateStatus(spacecraft.Status); err != nil {
		return err
	}

	return nil
}

func validateImage(image string) error {
	if _, err := url.Parse(image); err != nil {
		return ErrInvalidImage
	}

	return nil
}

func validateStatus(status Status) error {
	switch status {
	case StatusOperational:
		return nil
	case StatusDamaged:
		return nil
	}

	return ErrInvalidStatus
}
