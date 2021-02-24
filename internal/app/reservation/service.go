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
	customersRepo    CustomerRepository
}

// NewService creates new service instance.
func NewService(
	discountService app.DiscountService,
	bikeService app.BikeService,
	reservationsRepo Repository,
	customersRepo CustomerRepository,
) (*Service, error) {
	if discountService == nil {
		return nil, errors.New("empty discount service")
	}
	if bikeService == nil {
		return nil, errors.New("empty bike service")
	}
	if reservationsRepo == nil {
		return nil, errors.New("empty reservations repository")
	}
	if customersRepo == nil {
		return nil, errors.New("empty customers repository")
	}

	return &Service{
		discountService:  discountService,
		bikeService:      bikeService,
		reservationsRepo: reservationsRepo,
		customersRepo:    customersRepo,
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

	// Check if bike exists.
	if _, err := s.bikeService.Get(ctx, bikeID); err != nil {
		return false, fmt.Errorf("fetching bike data: %w", err)
	}

	reservations, err := s.reservationsRepo.ListReservations(ctx, ListReservationsQuery{
		BikeID:    bikeID,
		StartTime: startTime,
		EndTime:   endTime,
		Status:    app.ReservationStatusApproved,
		Limit:     1,
	})
	if err != nil {
		return false, fmt.Errorf("fetching reservations from repository: %w", err)
	}

	available := len(reservations) == 0

	return available, nil
}

// ListReservations returns list of reservations matching request criteria.
func (s *Service) ListReservations(ctx context.Context, req app.ListReservationsRequest) ([]app.Reservation, error) {
	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	reservations, err := s.reservationsRepo.ListReservations(ctx, ListReservationsQuery{
		BikeID:    req.BikeID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	})
	if err != nil {
		return nil, fmt.Errorf("fetching reservations from repository: %w", err)
	}
	return reservations, nil
}

// CreateReservation creates new reservation if possible.
// If creating reservation is not possible due to business logic or availability issues, this method returns valid response.
// If there are errors while processing request, returns nil and an error.
func (s *Service) CreateReservation(ctx context.Context, req app.CreateReservationRequest) (*app.ReservationResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	// We don't trust bike pricing from request,
	// so we fetch real bike data from bike service.
	bike, err := s.fetchRealBike(ctx, req.BikeID)
	if err != nil {
		if app.IsNotFoundError(err) {
			return &app.ReservationResponse{
				Status: app.ReservationStatusRejected,
				Reason: fmt.Sprintf("bike with id '%s' does not exists", req.BikeID),
			}, nil
		}
		return nil, err
	}

	// We don't trust bike pricing from request,
	// so we fetch real bike data from bike service.
	customer, err := s.updateCustomerData(ctx, req.Customer)
	if err != nil {
		return nil, err
	}

	value := s.calculateReservationValue(*bike, req.StartTime, req.EndTime)

	discountResp, err := s.discountService.CalculateDiscount(ctx, app.DiscountRequest{
		Customer:         customer,
		Location:         req.Location,
		Bike:             *bike,
		ReservationValue: value,
	})
	if err != nil {
		return nil, fmt.Errorf("checking available discounts: %w", err)
	}

	// We expect repository to return app.ConflictError if reservation for that bike in that time range already exists.
	reservation, err := s.reservationsRepo.CreateReservation(ctx, app.Reservation{
		ID:              uuid.New().String(),
		Status:          app.ReservationStatusApproved,
		Customer:        req.Customer,
		Bike:            *bike,
		StartTime:       req.StartTime,
		EndTime:         req.EndTime,
		TotalValue:      value - discountResp.Discount.Amount,
		AppliedDiscount: discountResp.Discount.Amount,
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
		Status:      reservation.Status,
		Reservation: reservation,
	}, nil
}

// CancelReservation removes reservation by id and bike id.
// Returns app.ErrNotFound if reservation doesn't exists.
func (s *Service) CancelReservation(ctx context.Context, bikeID string, id string) error {
	if err := s.reservationsRepo.CancelReservation(ctx, bikeID, id); err != nil {
		return fmt.Errorf("canceling reservation in repository: %w", err)
	}
	return nil
}

func (s *Service) fetchRealBike(ctx context.Context, bikeID string) (*app.Bike, error) {
	if bikeID == "" {
		return nil, errors.New("empty bike id")
	}

	existingBike, err := s.bikeService.Get(ctx, bikeID)
	if err != nil {
		return nil, fmt.Errorf("checking bike in repository: %w", err)
	}
	return existingBike, nil
}

func (s *Service) updateCustomerData(ctx context.Context, customer app.Customer) (app.Customer, error) {
	if customer.ID == "" {
		return customer, nil
	}

	existingCustomer, err := s.customersRepo.Get(ctx, customer.ID)
	if err != nil {
		return customer, fmt.Errorf("checking customer in repository: %w", err)
	}
	return *existingCustomer, nil
}

func (s *Service) calculateReservationValue(bike app.Bike, from, to time.Time) float64 {
	return bike.PricePerHour * to.Sub(from).Hours()
}
