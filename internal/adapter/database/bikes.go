package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/nglogic/go-example-project/internal/app"
)

// BikesRepository manages bikes in db.
type BikesRepository struct {
	db *sqlx.DB
}

// List returns list of all bikes from db sorted by name ascending.
func (r *BikesRepository) List(ctx context.Context) ([]app.Bike, error) {
	var bikes []bikeModel
	err := r.db.SelectContext(
		ctx,
		&bikes,
		`select * from bikes b order by b.model_name asc`,
	)
	if err != nil {
		return nil, fmt.Errorf("querying postgres: %w", err)
	}

	result := make([]app.Bike, 0, len(bikes))
	for _, b := range bikes {
		result = append(result, b.ToAppBike())
	}
	return result, nil
}

// Get returns a bike by id. If it doesn't exists, returns app.ErrNotFound error.
func (r *BikesRepository) Get(ctx context.Context, id string) (*app.Bike, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}

	var b bikeModel
	err := r.db.GetContext(
		ctx,
		&b,
		`select * from bikes where id = $1`,
		id,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, app.ErrNotFound
		}
		return nil, fmt.Errorf("querying postgres: %w", err)
	}

	result := b.ToAppBike()
	return &result, nil
}

// Add creates new bike in db.
func (r *BikesRepository) Add(ctx context.Context, b app.Bike) (id string, err error) {
	if b.ID == "" {
		b.ID = uuid.NewString()
	}

	_, err = r.db.NamedExecContext(
		ctx,
		`insert into bikes (id,model_name,weight,price_per_h)
		values (:id, :model_name, :weight, :price_per_h)`,
		newBikeModel(b),
	)
	if err != nil {
		return "", fmt.Errorf("inserting bike row into postgres: %w", err)
	}

	return b.ID, nil
}

// Update updates a bike in db by id. If bike is not in db, returns app.ErrNotFound error.
func (r *BikesRepository) Update(ctx context.Context, id string, b app.Bike) error {
	if id == "" {
		return fmt.Errorf("id is empty")
	}

	res, err := r.db.NamedExecContext(
		ctx,
		`update bikes set
			model_name=:model_name,
			weight=:weight,
			price_per_h=:price_per_h
		where id=:id`,
		newBikeModel(b),
	)
	if err != nil {
		return fmt.Errorf("inserting bike row into postgres: %w", err)
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return app.ErrNotFound
	}

	return nil
}

// Delete deletes a bike from db by id. If bike is not in db, returns app.ErrNotFound error.
func (r *BikesRepository) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is empty")
	}

	res, err := r.db.NamedExecContext(
		ctx,
		`delete from bikes where id=:id`,
		id,
	)
	if err != nil {
		return fmt.Errorf("deleting bike row from postgres: %w", err)
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return app.ErrNotFound
	}

	return nil
}

type bikeModel struct {
	ID           string  `db:"id"`
	ModelName    string  `db:"model_name"`
	Weight       float64 `db:"weight"`
	PricePerHour float64 `db:"price_per_h"`
}

func newBikeModel(ab app.Bike) bikeModel {
	// In this example bike is really simple and can be exaclty the same as domain bike type.
	return bikeModel(ab)
}

func (b *bikeModel) ToAppBike() app.Bike {
	return app.Bike(*b)
}
