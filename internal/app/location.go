package app

// Location represents physical location coordinates.
type Location struct {
	Lat float64
	Lng float64
}

// Validate validates location attributes.
func (l *Location) Validate() error {
	if l.Lat == 0 || l.Lng == 0 {
		return NewValidationError("invalid coordinates")
	}
	return nil
}
