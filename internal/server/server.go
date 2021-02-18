package server

import (
	context "context"
	"errors"

	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/nglogic/go-example-project/internal/app"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Server implements rpc ServiceServer.
type Server struct {
	bikeService        app.BikeService
	reservationService app.ReservationService
}

// New creates new Server instance.
func New(
	bikeService app.BikeService,
	reservationService app.ReservationService,
) (*Server, error) {
	if bikeService == nil {
		return nil, errors.New("bike service is nil")
	}
	if reservationService == nil {
		return nil, errors.New("reservation service is nil")
	}

	return &Server{
		bikeService:        bikeService,
		reservationService: reservationService,
	}, nil
}

// ListBikes returns list of all bikes.
func (s *Server) ListBikes(ctx context.Context, _ *empty.Empty) (*ListBikesResponse, error) {
	bikes, err := s.bikeService.List(ctx)
	if err != nil {
		return nil, NewGRPCError(err)
	}
	return newListBikesResponse(bikes), nil
}

// GetBike returns a bike.
func (s *Server) GetBike(ctx context.Context, req *GetBikeRequest) (*Bike, error) {
	b, err := s.bikeService.Get(ctx, req.Id)
	if err != nil {
		return nil, NewGRPCError(err)
	}
	return newResponseBike(b), nil
}

// CreateBike creates new bike.
func (s *Server) CreateBike(ctx context.Context, req *CreateBikeRequest) (*Bike, error) {
	if req.Bike == nil {
		return nil, status.Error(codes.InvalidArgument, "bike can't be empty")
	}
	b := newAppBikeFromRequest(req.Bike)
	createdBike, err := s.bikeService.Add(ctx, *b)
	if err != nil {
		return nil, NewGRPCError(err)
	}
	return newResponseBike(createdBike), nil
}

// DeleteBike deletes a bike.
func (s *Server) DeleteBike(ctx context.Context, req *DeleteBikeRequest) (*empty.Empty, error) {
	if err := s.bikeService.Delete(ctx, req.Id); err != nil {
		return nil, NewGRPCError(err)
	}

	return &empty.Empty{}, nil
}

// GetBikeAvailability checks bike availability in given time ranges.
func (s *Server) GetBikeAvailability(ctx context.Context, req *GetBikeAvailabilityRequest) (*GetBikeAvailabilityResponse, error) {
	available, err := s.reservationService.GetBikeAvailability(
		ctx,
		req.BikeId,
		req.StartTime.AsTime(),
		req.EndTime.AsTime(),
	)
	if err != nil {
		return nil, NewGRPCError(err)
	}

	return &GetBikeAvailabilityResponse{
		Available: available,
	}, nil
}
