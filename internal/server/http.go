package server

import (
	context "context"
	"fmt"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
)

// HTTP server default timeouts.
const (
	readTimeout       = 5 * time.Second
	readHeaderTimeout = 3 * time.Second
	writeTimeout      = 10 * time.Second
	idleTimeout       = 30 * time.Second
)

// RunHTTPServer starts http server with grpc gateway for ServiceServer.
// Server is gracefully shut down on context cancellation.
func RunHTTPServer(
	ctx context.Context,
	log logrus.FieldLogger,
	srv ServiceServer,
	addr string,
) error {
	mux := runtime.NewServeMux()

	err := RegisterServiceHandlerServer(ctx, mux, srv)
	if err != nil {
		return fmt.Errorf("registering http handlers for server: %w", err)
	}

	// See this great explanation on timeouts:
	// https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	s := http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadTimeout:       readTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
	}

	go func() {
		<-ctx.Done()
		log.Infof("http server: shutting down")
		if err := s.Shutdown(context.Background()); err != nil {
			log.Errorf("http server: failed to shutdown http gateway server: %v", err)
		}
	}()

	log.Infof("http server: listening on %s", addr)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}
