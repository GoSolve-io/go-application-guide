package discounts

import (
	"context"
	"errors"
	"fmt"

	"github.com/nglogic/go-example-project/internal/app"
)

const (
	// incidentsProximity defines size of a square in km, in which we search for incidents.
	incidentsProximity = 10.0
)

// Service provides methods for calculating discounts for a bike rental.
type Service struct {
	weatherService   app.WeatherService
	incidentsService app.BikeIncidentsService
}

// NewService creates new service instance.
func NewService(
	weather app.WeatherService,
	incidents app.BikeIncidentsService,
) (*Service, error) {
	if weather == nil {
		return nil, errors.New("empty weather service")
	}
	if incidents == nil {
		return nil, errors.New("empty incidents service")
	}

	return &Service{
		weatherService:   weather,
		incidentsService: incidents,
	}, nil
}

// CalculateDiscount returns available discount for a bike rental.
func (s *Service) CalculateDiscount(ctx context.Context, r *app.DiscountRequest) (*app.DiscountResponse, error) {
	if err := r.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	weather, err := s.weatherService.GetWeather(ctx, app.WeatherRequest{
		Location: r.Location,
	})
	if err != nil {
		return nil, fmt.Errorf("couldn't fetch weather data: %w", err)
	}

	incidents, err := s.incidentsService.GetIncidents(ctx, app.BikeIncidentsRequest{
		Location:  r.Location,
		Proximity: incidentsProximity,
	})
	if err != nil {
		return nil, fmt.Errorf("couldn't fetch incidents data: %w", err)
	}

	discount := selectOptimalDiscount(
		newBikeWeightDiscount(r.ReservationValue, r.Customer, r.Bike),
		newEnvironmentalDiscount(r.ReservationValue, r.Customer, weather, incidents),
		newBusinessCustomerDiscount(r.ReservationValue, r.Customer),
	)

	return &app.DiscountResponse{
		Discount: discount,
	}, nil
}
