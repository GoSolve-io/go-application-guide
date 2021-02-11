package bikes

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/nglogic/go-example-project/internal/app"
)

// Service provides methods for managing bikes for rental.
type Service struct {
	repository Repository
}

// NewService creates new service instance.
func NewService(bikeRepo Repository) (*Service, error) {
	if bikeRepo == nil {
		return nil, errors.New("empty bike repository")
	}
	return &Service{
		repository: bikeRepo,
	}, nil
}

// List returns all possible bikes.
func (s *Service) List(ctx context.Context) ([]app.Bike, error) {
	bikes, err := s.repository.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetching bikes from reposiytory: %w", err)
	}
	return bikes, nil
}

// Add adds a new bike.
func (s *Service) Add(ctx context.Context, b app.Bike) error {
	if b.ID != "" {
		return app.NewValidationError("can't add new bike with not empty id")
	}
	if err := b.Validate(); err != nil {
		return fmt.Errorf("invalid bike data: %w", err)
	}

	b.ID = uuid.New().String()
	if err := s.repository.Add(ctx, b); err != nil {
		return fmt.Errorf("adding bike to repository: %w", err)
	}
	return nil
}

// Update updates existing bike by id.
func (s *Service) Update(ctx context.Context, id string, b app.Bike) error {
	if id == "" {
		return app.NewValidationError("empty id")
	}
	if err := b.Validate(); err != nil {
		return fmt.Errorf("invalid bike data: %w", err)
	}

	exb, err := s.repository.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("fetching bike by id from repository: %w", err)
	}
	if exb == nil {
		return app.NewNotFoundError("bike not found")
	}

	if err := s.repository.Update(ctx, id, b); err != nil {
		return fmt.Errorf("updating bike in repository: %w", err)
	}
	return nil
}

// Delete deletes existing bike. If bike doesn't exists, returns nil.
func (s *Service) Delete(ctx context.Context, id string) error {
	if id == "" {
		return app.NewValidationError("empty id")
	}
	if err := s.repository.Delete(ctx, id); err != nil {
		if errors.Is(err, app.NotFoundError{}) {
			return nil
		}
		return fmt.Errorf("deleting bike in repository: %w", err)
	}
	return nil
}
