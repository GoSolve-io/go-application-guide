package metrics

import "errors"

var (
	// ErrInvalidTags means the provided tags are not valid
	ErrInvalidTags = errors.New("invalid tags")
	// ErrInvalidAttribute means the provided attribute is not valid
	ErrInvalidAttribute = errors.New("invalid attribute")
)
