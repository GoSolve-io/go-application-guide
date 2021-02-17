package server

import (
	context "context"
	"errors"

	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/nglogic/go-example-project/internal/app"
)

// Server implements rpc ServiceServer.
type Server struct {
	bikeService app.BikeService
}

// New creates new Server instance.
func New(
	bikeService app.BikeService,
) (*Server, error) {
	if bikeService == nil {
		return nil, errors.New("bike service is nil")
	}

	return &Server{
		bikeService: bikeService,
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
func (s *Server) GetBike(context.Context, *GetBikeRequest) (*Bike, error) {
	return nil, nil
}

// CreateBike creates new bike.
func (s *Server) CreateBike(context.Context, *CreateBikeRequest) (*Bike, error) {
	return nil, nil
}

// DeleteBike deletes a bike.
func (s *Server) DeleteBike(context.Context, *DeleteBikeRequest) (*empty.Empty, error) {
	return nil, nil
}
