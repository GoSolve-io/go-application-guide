package reservations

import (
	"context"

	"github.com/nglogic/go-example-project/internal/app"
)

// Repository provides methods for reading/writing reservation data.
type Repository interface {
	// CreateReservation creates new reservation for a bike.
	// If any reservation for this bike exists within given time range, will return app.ConflictError.
	CreateReservation(context.Context, app.Reservation) error
}
