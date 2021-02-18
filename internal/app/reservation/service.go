package reservation

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nglogic/go-example-project/internal/app"
)

// Service provides methods for making reservations.
type Service struct {
	discountService  app.DiscountService
	bikeService      app.BikeService
	reservationsRepo Repository
}

// NewService creates new service instance.
func NewService(discounts app.DiscountService, bikeService app.BikeService, reservationsRepo Repository) (*Service, error) {
	if discounts == nil {
		return nil, errors.New("empty discount service")
	}
	if bikeService == nil {
		return nil, errors.New("empty bike service")
	}
	if reservationsRepo == nil {
		return nil, errors.New("empty reservations repository")
	}

	return &Service{
		discountService:  discounts,
		bikeService:      bikeService,
		reservationsRepo: reservationsRepo,
	}, nil
}

// GetBikeAvailability returns true if bike with given id is available for rent in given time range.
func (s *Service) GetBikeAvailability(ctx context.Context, bikeID string, startTime, endTime time.Time) (bool, error) {
	if startTime.Before(time.Now()) {
		return false, app.NewValidationError("start time has to be in future")
	}
	if endTime.Before(startTime) {
		return false, app.NewValidationError("end time has to be after end time")
	}

	available, err := s.reservationsRepo.GetBikeAvailability(ctx, bikeID, startTime, endTime)
	if err != nil {
		return false, fmt.Errorf("checking bike availability in reservation repository: %w", err)
	}
	return available, nil
}

// MakeReservation creates new reservation if possible.
// If creating reservation is not possible due to business logic or availability issues, this method returns valid response.
// If there are errors while processing request, returns nil and an error.
func (s *Service) MakeReservation(ctx context.Context, req app.ReservationRequest) (*app.ReservationResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	// We don't trust bike pricing from request,
	// so we fetch real bike data from bike service.
	bike, err := s.fetchRealBike(ctx, req.Bike)
	if err != nil {
		return nil, err
	}
	if bike == nil {
		return &app.ReservationResponse{
			Status: app.ReservationStatusRejected,
			Reason: fmt.Sprintf("bike with id '%s' does not exists", req.Bike.ID),
		}, nil
	}

	value := s.calculateReservationValue(*bike, req.StartTime, req.EndTime)

	discountResp, err := s.discountService.CalculateDiscount(ctx, app.DiscountRequest{
		Customer:         req.Customer,
		Location:         req.Location,
		Bike:             *bike,
		ReservationValue: value,
	})
	if err != nil {
		return nil, fmt.Errorf("checking available discounts: %w", err)
	}

	// We expect repository to return app.ConflictError if reservation for that bike in that time range already exists.
	reservation, err := s.reservationsRepo.CreateReservation(ctx, app.Reservation{
		ID:         uuid.New().String(),
		Customer:   req.Customer,
		Bike:       req.Bike,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		TotalValue: value - discountResp.Discount.Amount,
	})
	if err != nil {
		if app.IsConflictError(err) {
			return &app.ReservationResponse{
				Status: app.ReservationStatusRejected,
				Reason: "bike not available in requested time range",
			}, nil
		}

		return nil, fmt.Errorf("creating reservation in repository: %w", err)
	}

	return &app.ReservationResponse{
		Status:                app.ReservationStatusApproved,
		Reservation:           reservation,
		AppliedDiscountAmount: discountResp.Discount.Amount,
	}, nil
}

func (s *Service) fetchRealBike(ctx context.Context, bike app.Bike) (*app.Bike, error) {
	if bike.ID == "" {
		return nil, errors.New("empty bike id")
	}

	existingBike, err := s.bikeService.Get(ctx, bike.ID)
	if err != nil {
		return nil, fmt.Errorf("checking bike in repository: %w", err)
	}
	return existingBike, nil
}

func (s *Service) calculateReservationValue(bike app.Bike, from, to time.Time) float64 {
	return bike.PricePerHour * to.Sub(from).Hours()
}
