package mysql

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/seniorescobar/alchemy-test/internal/domain/spacecraft"
)

const (
	table = "spacecrafts"

	colID        = "id"
	colName      = "name"
	colClass     = "class"
	colCrew      = "crew"
	colImage     = "image"
	colValue     = "value"
	colStatus    = "status"
	colArmaments = "armaments"
)

type SpacecraftRepository struct {
	db *sql.DB
}

func NewSpacecraftRepository(db *sql.DB) *SpacecraftRepository {
	return &SpacecraftRepository{
		db: db,
	}
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
	armaments, err := armamentsToJSON(spacecraft.Armaments)
	if err != nil {
		return fmt.Errorf("error marshaling armaments: %w", err)
	}

	query, args, err := sq.Insert(table).Columns(
		colName,
		colClass,
		colCrew,
		colImage,
		colValue,
		colStatus,
		colArmaments,
	).Values(
		spacecraft.Name,
		spacecraft.Class,
		spacecraft.Crew,
		spacecraft.Image,
		spacecraft.Value,
		spacecraft.Status,
		armaments,
	).ToSql()
	if err != nil {
		return err
	}

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	return nil
}

func (r *SpacecraftRepository) Update(ctx context.Context, spacecraft spacecraft.Spacecraft) error {
	return nil
}

func (r *SpacecraftRepository) Delete(ctx context.Context, id int) error {
	return nil
}

func armamentsToJSON(armaments []spacecraft.Armament) (json.RawMessage, error) {
	return json.Marshal(armaments)
}

func armamentsFromJSON(msg json.RawMessage) ([]spacecraft.Armament, error) {
	var armaments []spacecraft.Armament
	if err := json.Unmarshal(msg, &armaments); err != nil {
		return nil, err
	}

	return armaments, nil
}
