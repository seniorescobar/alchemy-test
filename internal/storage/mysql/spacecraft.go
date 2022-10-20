package mysql

import (
	"context"

	"github.com/seniorescobar/alchemy-test/internal/domain/spacecraft"
)

type SpacecraftRepository struct{}

func NewSpacecraftRepository() *SpacecraftRepository {
	return &SpacecraftRepository{}
}

func (r *SpacecraftRepository) List(ctx context.Context) ([]spacecraft.Spacecraft, error) {
	return []spacecraft.Spacecraft{
		{
			ID:   1,
			Name: "spacecraft 1",
		},
		{
			ID:   2,
			Name: "spacecraft 2",
		},
	}, nil
}

func (r *SpacecraftRepository) Get(ctx context.Context, id int) (spacecraft.Spacecraft, error) {
	return spacecraft.Spacecraft{
		ID:   id,
		Name: "spacecraft",
	}, nil
}

func (r *SpacecraftRepository) Create(ctx context.Context, spacecraft spacecraft.Spacecraft) error {
	return nil
}

func (r *SpacecraftRepository) Update(ctx context.Context, spacecraft spacecraft.Spacecraft) error {
	return nil
}

func (r *SpacecraftRepository) Delete(ctx context.Context, id int) error {
	return nil
}
