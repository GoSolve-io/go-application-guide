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
// Values are in fixed currency, we won't deal with currencies in decimals here for simplicity.
type Reservation struct {
	ID        string
	Customer  Customer
	Bike      Bike
	StartTime time.Time
	EndTime   time.Time

	// TotalValue is a total amount to pay by the customer.
	TotalValue float64

	// AppliedDiscount is amount of discount applied to total reservation value.
	AppliedDiscount float64
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
	GetBikeAvailability(ctx context.Context, bikeID string, startTime, endTime time.Time) (bool, error)
}

// ReservationRequest is a request for creating new reservation.
type ReservationRequest struct {
	BikeID    string
	Customer  Customer
	Location  Location
	StartTime time.Time
	EndTime   time.Time
}

// Validate validates request data.
func (r *ReservationRequest) Validate() error {
	if r.BikeID == "" {
		return NewValidationError("bike id is empty")
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
		return NewValidationError("start time time can't be in the past")
	}
	if r.EndTime.Before(r.StartTime) {
		return NewValidationError("end time have to ba after start time")
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
