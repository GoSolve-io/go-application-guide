package bikerental

import (
	"context"
	"fmt"
	"time"

	"github.com/nglogic/go-application-guide/internal/app"
)

// ReservationStatus describes reservation status.
type ReservationStatus string

// Reservation statuses.
const (
	ReservationStatusEmpty    ReservationStatus = ""
	ReservationStatusRejected ReservationStatus = "rejected"
	ReservationStatusApproved ReservationStatus = "approved"
	ReservationStatusCanceled ReservationStatus = "canceled"
)

// Reservation represents reservation for a bike.
// Values are in fixed currency, we won't deal with currencies in decimals here for simplicity.
type Reservation struct {
	ID        string
	Status    ReservationStatus
	Customer  Customer
	Bike      Bike
	StartTime time.Time
	EndTime   time.Time

	// TotalValue is a total amount to pay by the customer in eurocents.
	TotalValue int

	// AppliedDiscount is amount of discount applied to total reservation value in eurocents.
	AppliedDiscount int
}

// Validate validates reservation data.
func (r Reservation) Validate() error {
	if err := r.Customer.Validate(); err != nil {
		return fmt.Errorf("invalid customer data: %w", err)
	}

	return nil
}

// ReservationService provies methods for making reservations.
type ReservationService interface {
	GetBikeAvailability(ctx context.Context, bikeID string, startTime, endTime time.Time) (bool, error)
	ListReservations(ctx context.Context, req ListReservationsRequest) ([]Reservation, error)
	CreateReservation(ctx context.Context, req CreateReservationRequest) (*ReservationResponse, error)
	CancelReservation(ctx context.Context, bikeID string, id string) error
}

// CreateReservationRequest is a request for creating new reservation.
type CreateReservationRequest struct {
	BikeID    string
	Customer  Customer
	Location  Location
	StartTime time.Time
	EndTime   time.Time
}

// Validate validates request data.
func (r *CreateReservationRequest) Validate() error {
	if r.BikeID == "" {
		return app.NewValidationError("bike id is empty")
	}
	if r.Customer.ID == "" {
		if err := r.Customer.Validate(); err != nil {
			return fmt.Errorf("invalid customer data: %w", err)
		}
	}
	if err := r.Location.Validate(); err != nil {
		return fmt.Errorf("invalid location data: %w", err)
	}

	if r.StartTime.Before(time.Now()) {
		return app.NewValidationError("start time time can't be in the past")
	}
	if r.EndTime.Before(r.StartTime) {
		return app.NewValidationError("end time have to ba after start time")
	}

	return nil
}

// ReservationResponse is a response for create reservation request.
// If status is other than "approved", `Reservation` attribute will be nil.
type ReservationResponse struct {
	Status ReservationStatus

	// Reason contains reason of responding with given status.
	// If status is "approved", it should be empty.
	Reason string

	// Reservation will be empty for statuses other than "approved".
	Reservation *Reservation
}

// ListReservationsRequest is a request for listing reservations.
type ListReservationsRequest struct {
	BikeID    string
	StartTime time.Time
	EndTime   time.Time
}

// Validate validates request data.
func (r *ListReservationsRequest) Validate() error {
	if r.BikeID == "" {
		return app.NewValidationError("bike id can't be empty")
	}

	// Note: IsZero check doesn't work for empty timestamps created by empty protobuf timestamp.AsTime.
	if r.StartTime.Unix() == 0 {
		return app.NewValidationError("start time can't be empty")
	}
	if r.EndTime.Unix() == 0 {
		return app.NewValidationError("end time can't be empty")
	}
	if r.EndTime.Before(r.StartTime) {
		return app.NewValidationError("end time have to ba after start time")
	}

	return nil
}
