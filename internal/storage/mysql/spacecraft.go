package mysql

import (
	"context"

	"github.com/google/uuid"
	"github.com/seniorescobar/alchemy-test/internal/domain/spacecraft"
)

type SpacecraftRepository struct{}

func NewSpacecraftRepository() *SpacecraftRepository {
	return &SpacecraftRepository{}
}

func (r *SpacecraftRepository) List(ctx context.Context) ([]spacecraft.Spacecraft, error) {
	return []spacecraft.Spacecraft{
		{
			ID:   uuid.New(),
			Name: "spacecraft 1",
		},
		{
			ID:   uuid.New(),
			Name: "spacecraft 2",
		},
	}, nil
}

func (r *SpacecraftRepository) Get(ctx context.Context, id uuid.UUID) (spacecraft.Spacecraft, error) {
	return spacecraft.Spacecraft{
		ID:   id,
		Name: "spacecraft",
	}, nil
}
