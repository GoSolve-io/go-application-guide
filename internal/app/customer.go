package app

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
	Type      CustomerType
	FirstName string
	Surname   string
}

// Validate validates customer data.
func (c Customer) Validate() error {
	switch c.Type {
	case CustomerTypeIndividual, CustomerTypeBuisiness:
	default:
		return NewValidationError("invaid customer type")
	}

	if c.FirstName == "" {
		return NewValidationError("first name is required")
	}

	return nil
}
