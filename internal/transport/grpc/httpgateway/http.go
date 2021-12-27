package httpgateway

import (
	context "context"
	"errors"
	"fmt"
	"github.com/nglogic/go-application-guide/internal/adapter/metrics"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nglogic/go-application-guide/pkg/api/bikerentalv1"
	"github.com/sirupsen/logrus"
)

// HTTP server default timeouts.
const (
	readTimeout       = 5 * time.Second
	readHeaderTimeout = 3 * time.Second
	writeTimeout      = 10 * time.Second
	idleTimeout       = 30 * time.Second
)

// RunServer starts http server with grpc gateway for ServiceServer.
// Server is gracefully shut down on context cancellation.
func RunServer(
	ctx context.Context,
	log logrus.FieldLogger,
	met metrics.Provider,
	srv bikerentalv1.BikeRentalServiceServer,
	addr string,
) error {
	mux := runtime.NewServeMux()
	if err := bikerentalv1.RegisterBikeRentalServiceHandlerServer(ctx, mux, srv); err != nil {
		return fmt.Errorf("registering http handlers for server: %w", err)
	}

	var handler http.Handler = mux
	handler = HandlerWithLogCtx(handler)
	handler = HandlerWithTraceID(handler)
	handler = HandlerWithMetrics(handler, met)

	// See this great explanation on http timeouts:
	// https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	s := http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadTimeout:       readTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
	}

	go func() {
		<-ctx.Done()
		log.Info("http server: shutting down")
		if err := s.Shutdown(context.Background()); err != nil {
			log.Errorf("http server: failed to shutdown http gateway server: %v", err)
		}
	}()

	log.Infof("http server: listening on %s", addr)
	if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
