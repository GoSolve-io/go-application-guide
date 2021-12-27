package grpc

import (
	context "context"
	"time"

	"github.com/google/uuid"
	"github.com/nglogic/go-application-guide/internal/app"
	grpc "google.golang.org/grpc"
)

// TraceIDUnaryServerInterceptor returns a new unary server interceptor for generating trace id.
func TraceIDUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = app.CtxWithTraceID(ctx, uuid.NewString())

		return handler(ctx, req)
	}
}

// LogCtxUnaryServerInterceptor returns a new unary server interceptor adding request information to context for logging.
func LogCtxUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = app.CtxWithLogField(ctx, "grpc.method", info.FullMethod)

		return handler(ctx, req)
	}
}

// MetricsUnaryServerInterceptor returns a new unary server interceptor adding metrics to context for logging.
func MetricsUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		start := time.Now()
		defer func() {
			// TODO: use logger here, instead of adding this to context
			app.CtxWithLogField(ctx, "metrics_grpc_duration", time.Since(start).String())
		}()

		return handler(ctx, req)
	}
}
