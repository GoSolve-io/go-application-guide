package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/nglogic/go-example-project/internal/app"
	"github.com/sirupsen/logrus"
)

// Customer type enums for customer table.
const (
	customerTypeBusiness   = "business"
	customerTypeIndividual = "individual"
)

// ReservationsAdapter manages reservation data in db.
type ReservationsAdapter struct {
	db  *sqlx.DB
	log logrus.FieldLogger
}

// CreateReservation creates new reservation in db.
// Bike id must be provided.
// If customer doesn't exists, it is created with reservation.
func (a *ReservationsAdapter) CreateReservation(ctx context.Context, r app.Reservation) error {
	if err := a.checkReservationData(r); err != nil {
		return err
	}

	tx, err := a.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})
	if err != nil {
		return fmt.Errorf("creating postgresql transaction: %w", err)
	}
	defer rollbackTx(ctx, tx, a.log) // This will be noop after successful commit.

	if err := a.checkBikeExists(ctx, tx, r.Bike.ID); err != nil {
		return fmt.Errorf("invalid bike: %w", err)
	}

	if err := a.checkAvailability(ctx, tx, r); err != nil {
		return fmt.Errorf("bike not available: %w", err)
	}

	if r.Customer.ID != "" {
		if err := a.checkCustomerExists(ctx, tx, r.Customer.ID); err != nil {
			return fmt.Errorf("invalid customer: %w", err)
		}
	} else {
		r.Customer.ID = uuid.NewString()
		if err := a.createCustomer(ctx, tx, r.Customer); err != nil {
			return fmt.Errorf("creating customer: %w", err)
		}
	}

	if err := a.createReservation(ctx, tx, r); err != nil {
		return fmt.Errorf("creating reservation: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commiting postgres reservation transaction: %w", err)
	}

	return nil
}

func (a *ReservationsAdapter) checkReservationData(r app.Reservation) error {
	if r.ID == "" {
		return errors.New("reservation id is empty")
	}
	if r.Customer.ID == "" && r.Customer.Email == "" {
		return errors.New("customer id or email must be set")
	}
	if r.Bike.ID == "" {
		return errors.New("bike id is empty")
	}
	return nil
}

func (a *ReservationsAdapter) checkBikeExists(ctx context.Context, tx *sqlx.Tx, id string) error {
	err := tx.GetContext(
		ctx,
		nil,
		`SELECT id from bikes WHERE id=:id`,
		id,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return app.ErrNotFound
		}
		return fmt.Errorf("fetching bike by id from postgresql: %w", err)
	}
	return nil
}

func (a *ReservationsAdapter) checkCustomerExists(ctx context.Context, tx *sqlx.Tx, id string) error {
	err := tx.GetContext(
		ctx,
		nil,
		`SELECT id from customers WHERE id=:id`,
		id,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return app.ErrNotFound
		}
		return fmt.Errorf("fetching customer by id from postgresql: %w", err)
	}
	return nil
}

func (a *ReservationsAdapter) createCustomer(ctx context.Context, tx *sqlx.Tx, c app.Customer) error {
	m := newCustmerModel(c)
	_, err := a.db.NamedExecContext(
		ctx,
		`insert into customers (id, type, first_name, surname, email)
		values (:id, :type, :first_name, :surname, :email)`,
		m,
	)
	if err != nil {
		return fmt.Errorf("inserting customer row into postgres: %w", err)
	}
	return nil
}

func (a *ReservationsAdapter) checkAvailability(ctx context.Context, tx *sqlx.Tx, r app.Reservation) error {
	var count int
	m := newReservationModel(r)
	err := tx.GetContext(
		ctx,
		&count,
		`SELECT count(*) from reservations WHERE 
			bike_id = :bike_id
			AND from < :to
			AND to > :from
		`,
		m,
	)
	if err != nil {
		return fmt.Errorf("querying for conflicting reservations in postgresql: %w", err)
	}

	if count > 0 {
		return app.ErrConflict
	}
	return nil
}

func (a *ReservationsAdapter) createReservation(ctx context.Context, tx *sqlx.Tx, r app.Reservation) error {
	m := newReservationModel(r)
	_, err := a.db.NamedExecContext(
		ctx,
		`insert into reservations (id, bike_id, customer_id, from, to)
		values (:id, :bike_id, :customer_id, :from, :to)`,
		m,
	)
	if err != nil {
		return fmt.Errorf("inserting reservation row into postgres: %w", err)
	}
	return nil
}

type reservationModel struct {
	ID         string `db:"id"`
	BikeID     string `db:"bike_id"`
	CustomerID string `db:"customer_id"`
	From       string `db:"from"`
	To         string `db:"to"`
}

func newReservationModel(ar app.Reservation) reservationModel {
	return reservationModel{
		ID:         ar.ID,
		BikeID:     ar.Bike.ID,
		CustomerID: ar.Customer.ID,
		From:       ar.From.String(),
		To:         ar.To.String(),
	}
}

type customerModel struct {
	ID        string `db:"id"`
	Type      string `db:"type"`
	FirstName string `db:"first_name"`
	Surname   string `db:"surname"`
	Email     string `db:"email"`
}

func newCustmerModel(ac app.Customer) customerModel {
	c := customerModel{
		ID:        ac.ID,
		FirstName: ac.FirstName,
		Surname:   ac.Surname,
		Email:     ac.Email,
	}
	switch ac.Type {
	case app.CustomerTypeBuisiness:
		c.Type = customerTypeBusiness
	case app.CustomerTypeIndividual:
		c.Type = customerTypeIndividual
	}

	return c
}
