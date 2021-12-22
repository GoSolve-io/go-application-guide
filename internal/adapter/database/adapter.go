package database

import (
	"errors"
	"fmt"

	// Needed for registering postgres drivers in sql package.
	_ "github.com/lib/pq"

	// Needed for posgres migrations.
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/Masterminds/squirrel"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var sqlBuilder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

// Adapter is a storage adapter for app.
type Adapter struct {
	db  *sqlx.DB
	log logrus.FieldLogger
}

// NewAdapter creates new db adapter.
func NewAdapter(
	hostport string,
	dbname string,
	user string,
	pass string,
	migrationsDir string,
	log logrus.FieldLogger,
) (*Adapter, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, pass, hostport, dbname)

	m, err := migrate.New("file://"+migrationsDir, dbURL)
	if err != nil {
		return nil, fmt.Errorf("initiating migrations: %w", err)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, fmt.Errorf("migrating db: %w", err)
	}

	db, err := sqlx.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("opening postgres db: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("connection to postgres db failed: %w", err)
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
		db:  a.db,
		log: a.log.WithField("repository", "db.bikes"),
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
		db:  a.db,
		log: a.log.WithField("repository", "db.customers"),
	}
}
