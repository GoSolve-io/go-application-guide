package app

// ValidationError represents all kind of invalid data errors.
type ValidationError struct {
	Message string
}

// NewValidationError creates new ValidationError instance.
func NewValidationError(message string) error {
	return ValidationError{Message: message}
}

// Error fullfills error interface.
func (e ValidationError) Error() string {
	return e.Message
}

// NotFoundError represents all kind of errors resulting from not finding something.
type NotFoundError struct {
	Message string
}

// NewNotFoundError creates new NotFoundError instance.
func NewNotFoundError(message string) error {
	return NotFoundError{Message: message}
}

// Error fullfills error interface.
func (e NotFoundError) Error() string {
	return e.Message
}
