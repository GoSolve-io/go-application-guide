package reservations

import (
	"context"
	"time"

	"github.com/nglogic/go-example-project/internal/app"
)

// Repository provides methods for reading/writing reservation data.
type Repository interface {
	// StartTransaction creates new transaction.
	// Transaction should have at leat "repeatable reads" isolation level.
	StartTransaction(context.Context) (RepositoryTransaction, error)
}

// RepositoryTransaction represents a set of repository methods run in one transaction.
type RepositoryTransaction interface {
	ListReservations(bikeID string, from, to time.Time) ([]app.Reservation, error)
	CreateReservation(app.Reservation) error

	// Commit commits transaction. Should be ignored when Rollback is called first.
	Commit() error
	// Rollback rolls back transaction. Should be ignored when Commit was called first.
	Rollback() error
}
