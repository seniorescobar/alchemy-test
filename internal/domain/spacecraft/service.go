package spacecraft

import (
	"context"
	"net/url"
)

type (
	repo interface {
		List(context.Context, ...Filter) ([]Spacecraft, error)
		Get(context.Context, int) (Spacecraft, error)
		Create(context.Context, Spacecraft) error
		Update(context.Context, Spacecraft) error
		Delete(context.Context, int) error
	}

	Service struct {
		repo repo
	}

	Filter struct {
		Key   string
		Value string
	}
)

var (
	ErrInvalidImage  = &ValidationErr{"invalid image"}
	ErrInvalidStatus = &ValidationErr{"invalid status"}
	ErrInvalidFilter = &ValidationErr{"invalid filter"}

	AllowedFilters = []string{"name", "class", "status"}
)

func NewService(repo repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) List(ctx context.Context, filters ...Filter) ([]Spacecraft, error) {
	if err := validateFilters(filters...); err != nil {
		return nil, err
	}

	spacecrafts, err := s.repo.List(ctx, filters...)
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

func validateFilters(filters ...Filter) error {
	for _, filter := range filters {
		switch filter.Key {
		case "name":
			continue
		case "class":
			continue
		case "status":
			continue
		default:
			return ErrInvalidFilter
		}
	}

	return nil
}
