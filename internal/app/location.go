package app

import "fmt"

// Location represents physical location coordinates.
type Location struct {
	Lat  float64
	Long float64
}

// Validate validates location attributes.
func (l *Location) Validate() error {
	if l.Lat == 0 || l.Long == 0 {
		return NewValidationError("invalid coordinates")
	}
	return nil
}

// String returns string representation of location.
func (l *Location) String() string {
	return fmt.Sprintf("lat:%f long:%f", l.Lat, l.Long)
}
