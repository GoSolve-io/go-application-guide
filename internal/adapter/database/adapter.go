package database

import (
	"fmt"

	// Needed for registering postgres drivers in sql package.
	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

// Adapter is a storage adapter for app.
type Adapter struct {
	db  *sqlx.DB
	log logrus.FieldLogger
}

// NewAdapter creates new db adapter.
func NewAdapter(dsn string, log logrus.FieldLogger) (*Adapter, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("opening postgres db: %w", err)
	}

	return &Adapter{
		db:  db,
		log: log,
	}, nil
}

// Close closes underlying db client.
func (a *Adapter) Close() {
	a.db.Close()
}

// Bikes returns bikes repository.
func (a *Adapter) Bikes() *BikesRepository {
	return &BikesRepository{
		db: a.db,
	}
}

// Reservations returns reservations repository.
func (a *Adapter) Reservations() *ReservationsRepository {
	return &ReservationsRepository{
		parent: a,
		db:     a.db,
		log:    a.log.WithField("repository", "db.reservations"),
	}
}

// Customers returns customers repository.
func (a *Adapter) Customers() *CustomersRepository {
	return &CustomersRepository{
		db: a.db,
	}
}
