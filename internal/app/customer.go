package app

import (
	"fmt"

	"github.com/badoux/checkmail"
)

// CustomerType represents customer type.
type CustomerType int

// Customer types.
const (
	CustomerTypeUnknown CustomerType = iota
	CustomerTypeIndividual
	CustomerTypeBuisiness
)

// Customer represents customer renting a bike.
type Customer struct {
	ID        string
	Type      CustomerType
	FirstName string
	Surname   string
	Email     string
}

// Validate validates customer data.
func (c Customer) Validate() error {
	switch c.Type {
	case CustomerTypeIndividual, CustomerTypeBuisiness:
	default:
		return NewValidationError("invalid customer type")
	}

	if c.FirstName == "" {
		return NewValidationError("first name is required")
	}

	// For simplicity we use external library to validate emails.
	// We can live with it in this case, but in concequence our domain now relies on some external code!
	// So this is an exception that works only for well defined tasks (like email validation).
	if c.Email != "" {
		if err := checkmail.ValidateFormat(c.Email); err != nil {
			return ValidationError{
				Err: fmt.Errorf("invalid email address: %w", err),
			}
		}
	}

	return nil
}
