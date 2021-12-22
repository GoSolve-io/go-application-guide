package reservation

import (
	"context"
	"time"

	"github.com/nglogic/go-application-guide/internal/app/bikerental"
)

// Repository provides methods for reading/writing reservation data.
type Repository interface {
	// List returns list of reservations matching request criteria.
	List(context.Context, ListReservationsQuery) ([]bikerental.Reservation, error)

	// Get returns a reservation by id.
	// Returns app.ErrNotFound if reservation doesn't exist.
	Get(ctx context.Context, id string) (*bikerental.Reservation, error)

	// Create creates new reservation for a bike.
	// If any reservation for this bike exists within given time range, will return bikerental.ConflictError.
	// Returns created reservation data with filled all ids.
	Create(context.Context, bikerental.Reservation) (*bikerental.Reservation, error)

	// SetStatus updates the status of the reservation by its id.
	// Returns app.ErrNotFound if reservation doesn't exist.
	SetStatus(ctx context.Context, id string, status bikerental.ReservationStatus) error
}

// ListReservationsQuery is a set of filters for reservations result.
type ListReservationsQuery struct {
	BikeID    string
	StartTime time.Time
	EndTime   time.Time
	Status    bikerental.ReservationStatus
	Limit     int
}

// CustomerRepository provides methods for reading customer data.
type CustomerRepository interface {
	// Get returns customer by id.
	// Returns app.ErrNotFound if customer doesn't exist.
	Get(ctx context.Context, id string) (*bikerental.Customer, error)
}
