package server

import (
	context "context"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	grpc "google.golang.org/grpc"
)

// RunGRPCServer starts grpc server with ServiceServer service.
// Server is gracefully shut down on context cancellation.
func RunGRPCServer(
	ctx context.Context,
	log logrus.FieldLogger,
	srv ServiceServer,
	addr string,
) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("creating net listener: %w", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(TraceIDUnaryServerInterceptor()),
	)
	RegisterServiceServer(s, srv)
	go func() {
		<-ctx.Done()
		log.Infof("grpc server: shutting down")
		s.GracefulStop()
	}()

	log.Infof("grpc server: listening on %s", addr)
	return s.Serve(l)
}
