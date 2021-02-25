package bikerental

import (
	"fmt"

	"github.com/nglogic/go-example-project/internal/app"
)

// Location represents physical location coordinates.
type Location struct {
	Lat  float64
	Long float64
}

// Validate validates location attributes.
func (l *Location) Validate() error {
	if l.Lat == 0 || l.Long == 0 {
		return app.NewValidationError(fmt.Sprintf("invalid coordinates: %f, %f", l.Lat, l.Long))
	}
	return nil
}

// String returns string representation of location.
func (l *Location) String() string {
	return fmt.Sprintf("lat:%f long:%f", l.Lat, l.Long)
}
