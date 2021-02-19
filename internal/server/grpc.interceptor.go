package server

import (
	context "context"

	"github.com/google/uuid"
	"github.com/nglogic/go-example-project/internal/app"
	grpc "google.golang.org/grpc"
)

// TraceIDUnaryServerInterceptor returns a new unary server interceptor for generating trace id.
func TraceIDUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = app.CtxWithTraceID(ctx, uuid.NewString())
		return handler(ctx, req)
	}
}
