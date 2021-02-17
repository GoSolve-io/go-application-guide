package app

import "errors"

// Sentinel errors.
var (
	// ErrNotFound represents all kind of problems resulting from not finding something.
	ErrNotFound = errors.New("not found")

	// ErrConflict represents all kind of problems resulting from conflicting state.
	// For example - something was supposed to be created but it already exists.
	ErrConflict = errors.New("conflict")
)

// ValidationError represents all kind of invalid data errors.
type ValidationError struct {
	Err error
}

// NewValidationError creates new ValidationError instance.
func NewValidationError(message string) error {
	return ValidationError{Err: errors.New(message)}
}

// Error fullfills error interface.
func (e ValidationError) Error() string {
	return e.Err.Error()
}

// IsValidationError returns true if err has ValidationError in it's chain.
func IsValidationError(err error) bool {
	return errors.As(err, &ValidationError{})
}

// IsNotFoundError returns true if err has NotFoundError in it's chain.
func IsNotFoundError(err error) bool {
	return errors.Is(err, ErrNotFound)
}

// IsConflictError returns true if err has ConflictError in it's chain.
func IsConflictError(err error) bool {
	return errors.Is(err, ErrConflict)
}
