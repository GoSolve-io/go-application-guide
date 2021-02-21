package httpgateway

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/nglogic/go-example-project/internal/app"
)

// HandlerWithTraceID wraps handler with middleware generating trace id for each request.
func HandlerWithTraceID(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := app.CtxWithTraceID(r.Context(), uuid.NewString())
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}

// HandlerWithLogCtx wraps handler with middleware adding request information to context for logging.
func HandlerWithLogCtx(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = app.CtxWithLogField(ctx, "http.url", r.URL.String())
		ctx = app.CtxWithLogField(ctx, "http.method", r.Method)

		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}
