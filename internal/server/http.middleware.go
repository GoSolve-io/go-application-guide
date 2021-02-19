package server

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/nglogic/go-example-project/internal/app"
)

// HTTPHandlerWithTraceID wraps handler with middleware generating trace id for each request.
func HTTPHandlerWithTraceID(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := app.CtxWithTraceID(r.Context(), uuid.NewString())
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}
