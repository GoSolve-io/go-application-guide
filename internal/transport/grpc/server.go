package grpc

import (
	context "context"
	"errors"
	"fmt"
	"net"

	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/nglogic/go-example-project/internal/app"
	v1 "github.com/nglogic/go-example-project/pkg/api/v1"
	"github.com/sirupsen/logrus"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Server implements rpc ServiceServer.
type Server struct {
	bikeService        app.BikeService
	reservationService app.ReservationService
	log                logrus.FieldLogger
}

// NewServer creates new Server instance.
func NewServer(
	bikeService app.BikeService,
	reservationService app.ReservationService,
	log logrus.FieldLogger,
) (*Server, error) {
	if bikeService == nil {
		return nil, errors.New("bike service is nil")
	}
	if reservationService == nil {
		return nil, errors.New("reservation service is nil")
	}
	if log == nil {
		return nil, errors.New("logger is nil")
	}

	return &Server{
		bikeService:        bikeService,
		reservationService: reservationService,
		log:                log,
	}, nil
}

// ListBikes returns list of all bikes.
func (s *Server) ListBikes(ctx context.Context, _ *empty.Empty) (*v1.ListBikesResponse, error) {
	bikes, err := s.bikeService.List(ctx)
	if err != nil {
		s.logError(ctx, err, "ListBikes")
		return nil, NewServerError(err)
	}
	return newListBikesResponse(bikes), nil
}

// GetBike returns a bike.
func (s *Server) GetBike(ctx context.Context, req *v1.GetBikeRequest) (*v1.Bike, error) {
	b, err := s.bikeService.Get(ctx, req.Id)
	if err != nil {
		s.logError(ctx, err, "GetBike")
		return nil, NewServerError(err)
	}
	return newResponseBike(b), nil
}

// CreateBike creates new bike.
func (s *Server) CreateBike(ctx context.Context, req *v1.CreateBikeRequest) (*v1.Bike, error) {
	if req.Bike == nil {
		return nil, status.Error(codes.InvalidArgument, "bike can't be empty")
	}
	b := newAppBikeFromRequest(req.Bike)
	createdBike, err := s.bikeService.Add(ctx, *b)
	if err != nil {
		s.logError(ctx, err, "CreateBike")
		return nil, NewServerError(err)
	}

	s.logInfo(ctx, "CreateBike", "bike created: %s", createdBike.ID)

	return newResponseBike(createdBike), nil
}

// DeleteBike deletes a bike.
func (s *Server) DeleteBike(ctx context.Context, req *v1.DeleteBikeRequest) (*empty.Empty, error) {
	if err := s.bikeService.Delete(ctx, req.Id); err != nil {
		s.logError(ctx, err, "DeleteBike")
		return nil, NewServerError(err)
	}

	s.logInfo(ctx, "DeleteBike", "bike delete ok: %s", req.Id)

	return &empty.Empty{}, nil
}

// GetBikeAvailability checks bike availability in given time ranges.
func (s *Server) GetBikeAvailability(ctx context.Context, req *v1.GetBikeAvailabilityRequest) (*v1.GetBikeAvailabilityResponse, error) {
	available, err := s.reservationService.GetBikeAvailability(
		ctx,
		req.BikeId,
		req.StartTime.AsTime(),
		req.EndTime.AsTime(),
	)
	if err != nil {
		s.logError(ctx, err, "GetBikeAvailability")
		return nil, NewServerError(err)
	}

	return &v1.GetBikeAvailabilityResponse{
		Available: available,
	}, nil
}

// ListReservations returns list of reservations for a bike.
func (s *Server) ListReservations(ctx context.Context, req *v1.ListReservationsRequest) (*v1.ListReservationsResponse, error) {
	reservations, err := s.reservationService.ListReservations(ctx, app.ListReservationsRequest{
		BikeID:    req.BikeId,
		StartTime: req.StartTime.AsTime(),
		EndTime:   req.EndTime.AsTime(),
	})
	if err != nil {
		s.logError(ctx, err, "ListReservations")
		return nil, NewServerError(err)
	}

	var outrs []*v1.Reservation
	for _, v := range reservations {
		outrs = append(outrs, newResponseReservation(&v))
	}
	return &v1.ListReservationsResponse{
		Reservations: outrs,
	}, nil
}

// CreateReservation creates new reservation.
// Returns created object with new id.
func (s *Server) CreateReservation(ctx context.Context, req *v1.CreateReservationRequest) (*v1.CreateReservationResponse, error) {
	if req.Customer == nil {
		return nil, status.Error(codes.InvalidArgument, "customer can't be empty")
	}
	customer := newAppCustomerFromRequest(req.Customer)

	if req.Location == nil {
		return nil, status.Error(codes.InvalidArgument, "location can't be empty")
	}
	location := newAppLocationFromRequest(req.Location)

	resp, err := s.reservationService.CreateReservation(ctx, app.CreateReservationRequest{
		BikeID:    req.BikeId,
		Customer:  *customer,
		Location:  *location,
		StartTime: req.StartTime.AsTime(),
		EndTime:   req.EndTime.AsTime(),
	})
	if err != nil {
		s.logError(ctx, err, "CreateReservation")
		return nil, NewServerError(err)
	}

	if resp.Reservation != nil {
		s.logInfo(ctx, "CreateReservation", "reservation created: %s", resp.Reservation.ID)
	} else {
		s.logInfo(ctx, "CreateReservation", "reservation not created, reason: %s", resp.Reason)
	}

	return newCreateReservationResponse(resp), nil
}

// CancelReservation cancels reservation for a bike.
func (s *Server) CancelReservation(ctx context.Context, req *v1.CancelReservationRequest) (*empty.Empty, error) {
	if err := s.reservationService.CancelReservation(ctx, req.BikeId, req.Id); err != nil {
		s.logError(ctx, err, "CancelReservation")
		return nil, NewServerError(err)
	}

	return &empty.Empty{}, nil
}

func (s *Server) logError(ctx context.Context, err error, endpoint string) {
	// Don't log errors caused by invalid user input.
	if app.IsValidationError(err) {
		return
	}

	app.AugmentLogFromCtx(ctx, s.log).Errorf("handling request for %s: %v", endpoint, err)
}

func (s *Server) logInfo(ctx context.Context, endpoint string, format string, args ...interface{}) {
	app.AugmentLogFromCtx(ctx, s.log).Infof(
		fmt.Sprintf("handling request for %s: %s", endpoint, format),
		args...,
	)
}

// RunServer starts grpc server with ServiceServer service.
// Server is gracefully shut down on context cancellation.
func RunServer(
	ctx context.Context,
	log logrus.FieldLogger,
	srv v1.ServiceServer,
	addr string,
) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("creating net listener: %w", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(TraceIDUnaryServerInterceptor()),
	)
	v1.RegisterServiceServer(s, srv)
	go func() {
		<-ctx.Done()
		log.Infof("grpc server: shutting down")
		s.GracefulStop()
	}()

	log.Infof("grpc server: listening on %s", addr)
	return s.Serve(l)
}
