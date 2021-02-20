package reservation

import (
	"context"
	"time"

	"github.com/nglogic/go-example-project/internal/app"
)

// Repository provides methods for reading/writing reservation data.
type Repository interface {
	// GetBikeAvailability returns true if bike with given id is available for rent in given time range.
	GetBikeAvailability(ctx context.Context, bikeID string, startTime, endTime time.Time) (bool, error)

	// ListReservations returns list of reservations matching request criteria.
	ListReservations(context.Context, app.ListReservationsRequest) ([]app.Reservation, error)

	// CreateReservation creates new reservation for a bike.
	// If any reservation for this bike exists within given time range, will return app.ConflictError.
	// Returns created reservation data with filled all ids.
	CreateReservation(context.Context, app.Reservation) (*app.Reservation, error)

	// CancelReservation cancels reservation by id and bike id.
	// Canceled reservation is not deleted, but gets status "canceled".
	// Returns app.ErrNotFound if reservation doesn't exists.
	CancelReservation(ctx context.Context, bikeID string, id string) error
}

// CustomerRepository provides methods for reading customer data.
type CustomerRepository interface {
	// Get returns customer by id.
	// Returns app.ErrNotFound if customer doesn't exist.
	Get(ctx context.Context, id string) (*app.Customer, error)
}
