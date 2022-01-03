package grpc

import (
	context "context"
	"time"

	"github.com/nglogic/go-application-guide/internal/adapter/metrics"

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
func MetricsUnaryServerInterceptor(m metrics.Provider) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()

		m.Count("grpc.invocation", info.FullMethod)

		resp, err := handler(ctx, req)
		if err != nil {
			m.Count("grpc.errors_count", info.FullMethod)
		}

		m.Duration(time.Since(start), info.FullMethod)

		return resp, err
	}
}
