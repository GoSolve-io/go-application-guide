package app

import "errors"

// internalError is a base type for all app errors.
// This is a helper type to not repeat Error() functions.
type internalError struct {
	Err error
}

func newInternalErrorFromString(message string) internalError {
	return internalError{Err: errors.New(message)}
}

// Error fullfills error interface.
func (e internalError) Error() string {
	return e.Err.Error()
}

// ValidationError represents all kind of invalid data errors.
type ValidationError struct {
	internalError
}

// NewValidationError creates new ValidationError instance.
func NewValidationError(message string) error {
	return ValidationError{newInternalErrorFromString(message)}
}

// IsValidationError returns true if err has ValidationError in it's chain.
func IsValidationError(err error) bool {
	return errors.As(err, &ValidationError{})
}

// NotFoundError represents all kind of errors resulting from not finding something.
type NotFoundError struct {
	internalError
}

// NewNotFoundError creates new NotFoundError instance.
func NewNotFoundError(message string) error {
	return NotFoundError{newInternalErrorFromString(message)}
}

// IsNotFoundError returns true if err has NotFoundError in it's chain.
func IsNotFoundError(err error) bool {
	return errors.As(err, &NotFoundError{})
}

// ConflictError represents all kind of errors resulting from conflicting state.
// For example - something was supposed to be created but it already exists.
type ConflictError struct {
	internalError
}

// NewConflictError creates new ConflictError instance.
func NewConflictError(message string) error {
	return ConflictError{newInternalErrorFromString(message)}
}

// IsConflictError returns true if err has ConflictError in it's chain.
func IsConflictError(err error) bool {
	return errors.As(err, &ConflictError{})
}
