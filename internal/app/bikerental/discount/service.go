package discount

import (
	"context"
	"errors"
	"fmt"

	"github.com/nglogic/go-application-guide/internal/app"
	"github.com/nglogic/go-application-guide/internal/app/bikerental"
)

const (
	// incidentsProximity defines size of a square in km, in which we search for incidents.
	incidentsProximity = 10.0
)

// Service provides methods for calculating discounts for a bike rental.
type Service struct {
	weatherService   bikerental.WeatherService
	incidentsService bikerental.BikeIncidentsService
}

// NewService creates new service instance.
func NewService(
	weather bikerental.WeatherService,
	incidents bikerental.BikeIncidentsService,
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
func (s *Service) CalculateDiscount(ctx context.Context, r bikerental.DiscountRequest) (*bikerental.DiscountResponse, error) {
	if err := r.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	weather, err := s.weatherService.GetWeather(ctx, bikerental.WeatherRequest{
		Location: r.Location,
	})
	if err != nil {
		// We're ok with nil weather value if weather for given location is not found.
		if app.IsNotFoundError(err) {
			weather = nil
		} else {
			return nil, fmt.Errorf("couldn't fetch weather data: %w", err)
		}
	}

	incidents, err := s.incidentsService.GetIncidents(ctx, bikerental.BikeIncidentsRequest{
		Location:  r.Location,
		Proximity: incidentsProximity,
	})
	if err != nil {
		// We're ok with nil incidents info if data for given location is not found.
		if app.IsNotFoundError(err) {
			incidents = nil
		} else {
			return nil, fmt.Errorf("couldn't fetch incidents data: %w", err)
		}
	}

	discount := selectOptimalDiscount(
		newBikeWeightDiscount(r.ReservationValue, r.Customer, r.Bike),
		newTemperatureDiscount(r.ReservationValue, r.Customer, weather),
		newIncidentsDiscount(r.ReservationValue, r.Customer, incidents),
		newBusinessCustomerDiscount(r.ReservationValue, r.Customer),
	)

	return &bikerental.DiscountResponse{
		Discount: discount,
	}, nil
}
