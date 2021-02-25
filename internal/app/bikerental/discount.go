package bikerental

import (
	"context"
	"fmt"

	"github.com/nglogic/go-example-project/internal/app"
)

// Discount represents fixed discount for bike rental.
type Discount struct {
	// Amount is in euros, we wont deal with currencies in decimals here for simplicity.
	Amount float64
	Reason string
}

// DiscountService provides methods for calculating discounts for a bike rentals.
type DiscountService interface {
	CalculateDiscount(context.Context, DiscountRequest) (*DiscountResponse, error)
}

// DiscountRequest is a request for determining a discount for a bike rental.
type DiscountRequest struct {
	Customer         Customer
	Location         Location
	Bike             Bike
	ReservationValue float64
}

// Validate validates the request.
func (r DiscountRequest) Validate() error {
	if err := r.Customer.Validate(); err != nil {
		return fmt.Errorf("invalid customer data: %w", err)
	}

	if r.Location.Lat == 0 || r.Location.Long == 0 {
		return app.NewValidationError("invalid location")
	}

	if r.ReservationValue == 0 {
		return app.NewValidationError("empty reservation value")
	}

	if r.Bike.Weight == 0.0 {
		return app.NewValidationError("empty bike weight")
	}

	return nil
}

// DiscountResponse is a reponse with calculated discount.
type DiscountResponse struct {
	Discount Discount
}
