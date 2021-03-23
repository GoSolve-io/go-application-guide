package reservation

import (
	"context"
	"time"

	"github.com/nglogic/go-application-guide/internal/app/bikerental"
)

// Repository provides methods for reading/writing reservation data.
type Repository interface {
	// ListReservations returns list of reservations matching request criteria.
	ListReservations(context.Context, ListReservationsQuery) ([]bikerental.Reservation, error)

	// CreateReservation creates new reservation for a bike.
	// If any reservation for this bike exists within given time range, will return bikerental.ConflictError.
	// Returns created reservation data with filled all ids.
	CreateReservation(context.Context, bikerental.Reservation) (*bikerental.Reservation, error)

	// CancelReservation cancels reservation by id and bike id.
	// Canceled reservation is not deleted, but gets status "canceled".
	// Returns bikerental.ErrNotFound if reservation doesn't exists.
	CancelReservation(ctx context.Context, bikeID string, id string) error
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
	// Returns bikerental.ErrNotFound if customer doesn't exist.
	Get(ctx context.Context, id string) (*bikerental.Customer, error)
}
