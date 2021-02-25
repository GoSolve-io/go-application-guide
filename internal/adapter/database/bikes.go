package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/nglogic/go-example-project/internal/app"
	"github.com/nglogic/go-example-project/internal/app/bikerental"
	"github.com/sirupsen/logrus"
)

// BikesRepository manages bikes in db.
type BikesRepository struct {
	db  *sqlx.DB
	log logrus.FieldLogger
}

// List returns list of all bikes from db sorted by name ascending.
func (r *BikesRepository) List(ctx context.Context) ([]bikerental.Bike, error) {
	var bikes []bikeModel
	if err := r.db.SelectContext(ctx, &bikes, "select * from bikes order by model_name asc"); err != nil {
		return nil, fmt.Errorf("querying postgres: %w", err)
	}

	result := make([]bikerental.Bike, 0, len(bikes))
	for _, b := range bikes {
		result = append(result, b.ToAppBike())
	}
	return result, nil
}

// Get returns a bike by id. If it doesn't exists, returns bikerental.ErrNotFound error.
func (r *BikesRepository) Get(ctx context.Context, id string) (*bikerental.Bike, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}

	var b bikeModel
	if err := r.db.GetContext(ctx, &b, "select * from bikes where id=$1", id); err != nil {
		if err == sql.ErrNoRows {
			return nil, app.ErrNotFound
		}
		return nil, fmt.Errorf("querying postgres: %w", err)
	}

	result := b.ToAppBike()
	return &result, nil
}

// Add creates new bike in db.
func (r *BikesRepository) Add(ctx context.Context, b bikerental.Bike) (id string, err error) {
	if b.ID == "" {
		b.ID = uuid.NewString()
	}

	sqlq := sqlBuilder.Insert("bikes").
		Columns("id", "model_name", "weight", "price_per_h").
		Values(
			squirrel.Expr(":id"),
			squirrel.Expr(":model_name"),
			squirrel.Expr(":weight"),
			squirrel.Expr(":price_per_h"),
		)
	q, _, err := sqlq.ToSql()
	if err != nil {
		return "", fmt.Errorf("building sql query: %w", err)
	}

	if _, err = r.db.NamedExecContext(ctx, q, newBikeModel(b)); err != nil {
		return "", fmt.Errorf("inserting bike row into postgres: %w", err)
	}

	app.AugmentLogFromCtx(ctx, r.log).WithField("id", b.ID).Info("bike created in db")

	return b.ID, nil
}

// Update updates a bike in db by id. If bike is not in db, returns bikerental.ErrNotFound error.
func (r *BikesRepository) Update(ctx context.Context, id string, b bikerental.Bike) error {
	if id == "" {
		return fmt.Errorf("id is empty")
	}

	sqlq := sqlBuilder.Update("bikes").
		Set("model_name", squirrel.Expr(":model_name")).
		Set("weight", squirrel.Expr(":weight")).
		Set("price_per_h", squirrel.Expr(":price_per_h")).
		Where(squirrel.Eq{"id": squirrel.Expr(":id")})
	q, _, err := sqlq.ToSql()
	if err != nil {
		return fmt.Errorf("building sql query: %w", err)
	}

	res, err := r.db.NamedExecContext(ctx, q, newBikeModel(b))
	if err != nil {
		return fmt.Errorf("inserting bike row into postgres: %w", err)
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return app.ErrNotFound
	}

	app.AugmentLogFromCtx(ctx, r.log).WithField("id", id).Info("bike updated in db")

	return nil
}

// Delete deletes a bike from db by id. If bike is not in db, returns bikerental.ErrNotFound error.
func (r *BikesRepository) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is empty")
	}

	res, err := r.db.ExecContext(ctx, `delete from bikes where id=$1`, id)
	if err != nil {
		return fmt.Errorf("deleting bike row from postgres: %w", err)
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return app.ErrNotFound
	}

	app.AugmentLogFromCtx(ctx, r.log).WithField("id", id).Info("bike deleted from db")

	return nil
}

type bikeModel struct {
	ID           string  `db:"id"`
	ModelName    string  `db:"model_name"`
	Weight       float64 `db:"weight"`
	PricePerHour float64 `db:"price_per_h"`
}

func newBikeModel(ab bikerental.Bike) bikeModel {
	// In this example bike is really simple and can be exaclty the same as domain bike type.
	return bikeModel(ab)
}

func (b *bikeModel) ToAppBike() bikerental.Bike {
	return bikerental.Bike(*b)
}
