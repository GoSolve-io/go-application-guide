package server

import (
	"github.com/nglogic/go-example-project/internal/app"
	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// NewGRPCError creates error for grp response.
func NewGRPCError(err error) error {
	if err == nil {
		return nil
	}

	var code codes.Code
	switch {
	case app.IsNotFoundError(err):
		code = codes.NotFound
	case app.IsConflictError(err):
		code = codes.AlreadyExists
	case app.IsValidationError(err):
		code = codes.InvalidArgument
	default:
		code = codes.Internal
	}

	return status.Error(code, err.Error())
}
