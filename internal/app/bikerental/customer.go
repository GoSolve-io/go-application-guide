package bikerental

import (
	"fmt"

	"github.com/badoux/checkmail"
	"github.com/nglogic/go-application-guide/internal/app"
)

// CustomerType represents customer type.
type CustomerType int

// Customer types.
const (
	CustomerTypeUnknown CustomerType = iota
	CustomerTypeIndividual
	CustomerTypeBusiness
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
	case CustomerTypeIndividual, CustomerTypeBusiness:
	default:
		return app.NewValidationError("invalid customer type")
	}

	if c.FirstName == "" {
		return app.NewValidationError("first name is required")
	}

	// For simplicity, we use external library to validate emails.
	// We can live with it in this case, but in consequence our domain now relies on some external code!
	// So this is an exception that works only for well-defined tasks (like email validation).
	if c.Email != "" {
		if err := checkmail.ValidateFormat(c.Email); err != nil {
			return app.ValidationError{
				Err: fmt.Errorf("invalid email address: %w", err),
			}
		}
	}

	return nil
}
