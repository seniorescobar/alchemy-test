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

func (r *SpacecraftRepository) List(ctx context.Context, filters ...spacecraft.Filter) ([]spacecraft.Spacecraft, error) {
	sb := sq.
		Select(
			colID,
			colName,
			colCrew,
			colImage,
			colValue,
			colStatus,
			colArmaments,
		).
		From(table)

	for _, filter := range filters {
		sb = sb.Where(sq.Eq{filter.Key: filter.Value})
	}

	query, args, err := sb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building query: %w", err)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error querying rows: %w", err)
	}

	scs := make([]spacecraft.Spacecraft, 0)
	for rows.Next() {
		var sc Spacecraft
		if err := rows.Scan(
			&sc.ID,
			&sc.Name,
			&sc.Crew,
			&sc.Image,
			&sc.Val,
			&sc.Status,
			&sc.Armaments,
		); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		armaments, err := armamentsFromJSON(sc.Armaments)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling armaments: %w", err)
		}

		sc.Spacecraft.Armaments = armaments
		scs = append(scs, sc.Spacecraft)
	}

	return scs, nil
}

func (r *SpacecraftRepository) Get(ctx context.Context, id int) (spacecraft.Spacecraft, error) {
	query, args, err := sq.Select(
		colID,
		colName,
		colCrew,
		colImage,
		colValue,
		colStatus,
		colArmaments,
	).
		From(table).
		Where(sq.Eq{colID: id}).
		ToSql()
	if err != nil {
		return spacecraft.Spacecraft{}, fmt.Errorf("error building query: %w", err)
	}

	var sc Spacecraft
	if err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&sc.ID,
		&sc.Name,
		&sc.Crew,
		&sc.Image,
		&sc.Val,
		&sc.Status,
		&sc.Armaments,
	); err != nil {
		return spacecraft.Spacecraft{}, fmt.Errorf("error scanning row: %w", err)
	}

	armaments, err := armamentsFromJSON(sc.Armaments)
	if err != nil {
		return spacecraft.Spacecraft{}, fmt.Errorf("error unmarshaling armaments: %w", err)
	}

	sc.Spacecraft.Armaments = armaments
	return sc.Spacecraft, nil
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
		spacecraft.Val,
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

func (r *SpacecraftRepository) Update(ctx context.Context, sc spacecraft.Spacecraft) error {
	armaments, err := armamentsToJSON(sc.Armaments)
	if err != nil {
		return fmt.Errorf("error marshaling armaments: %w", err)
	}

	query, args, err := sq.
		Update(table).
		SetMap(map[string]interface{}{
			colName:      sc.Name,
			colClass:     sc.Class,
			colCrew:      sc.Crew,
			colImage:     sc.Image,
			colValue:     sc.Val,
			colStatus:    sc.Status,
			colArmaments: armaments,
		}).
		Where(sq.Eq{colID: sc.ID}).
		ToSql()
	if err != nil {
		return err
	}

	res, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if ra == 0 {
		return spacecraft.ErrSpacecraftNotFound
	}

	return nil
}

func (r *SpacecraftRepository) Delete(ctx context.Context, id int) error {
	query, args, err := sq.Delete(table).Where(sq.Eq{colID: id}).ToSql()
	if err != nil {
		return fmt.Errorf("error building query: %w", err)
	}

	res, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if ra == 0 {
		return spacecraft.ErrSpacecraftNotFound
	}

	return nil
}

type Spacecraft struct {
	spacecraft.Spacecraft

	Armaments json.RawMessage
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
