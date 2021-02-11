package app

import (
	"context"
	"fmt"
	"time"
)

// ReservationStatus describes reservation status.
type ReservationStatus string

// Reservation statuses.
const (
	ReservationStatusRejected ReservationStatus = "rejected"
	ReservationStatusApproved ReservationStatus = "approved"
)

// Reservation represents reservation for a bike.
type Reservation struct {
	ID       string
	Customer Customer
	Bike     Bike
	From     time.Time
	To       time.Time
	// TotalValue is in euros, we won't deal with currencies in decimals here for simplicity.
	TotalValue float64
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
	MakeReservation(context.Context, ReservationRequest) (*ReservationResponse, error)
}

// ReservationRequest is a request for creating new reservation.
type ReservationRequest struct {
	Customer Customer
	Bike     Bike
	Location Location
	From     time.Time
	To       time.Time
}

// Validate validates request data.
func (r *ReservationRequest) Validate() error {
	if err := r.Customer.Validate(); err != nil {
		return fmt.Errorf("invalid customer data: %w", err)
	}
	if err := r.Bike.Validate(); err != nil {
		return fmt.Errorf("invalid bike data: %w", err)
	}
	if err := r.Location.Validate(); err != nil {
		return fmt.Errorf("invalid location data: %w", err)
	}

	if r.From.Before(time.Now()) {
		return NewValidationError("'from' time can't be in the past")
	}
	if r.To.Before(r.From) {
		return NewValidationError("'to' have to ba after 'from'")
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

	// AppliedDiscountAmount is amount of discount applied to total reservation value.
	AppliedDiscountAmount float64
}
