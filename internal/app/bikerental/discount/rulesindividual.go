package discount

// This file contains all business rules for calculating discounts.

import (
	"math"

	"github.com/nglogic/go-application-guide/internal/app/bikerental"
)

// newBikeWeightDiscount returns discount for individual customers based on reservation value and bike weight.
// Discount rules:
// - individual customers only
// - bike weight >= 15kg
// - maximum discount is 20% of reservation value.
func newBikeWeightDiscount(resValue int, customer bikerental.Customer, bike bikerental.Bike) bikerental.Discount {
	if customer.Type != bikerental.CustomerTypeIndividual {
		return bikerental.Discount{}
	}
	if bike.Weight < 15 {
		return bikerental.Discount{}
	}

	discountPercent := bike.Weight - 15.0
	if discountPercent > 20 {
		discountPercent = 20
	}

	return bikerental.Discount{
		Amount: int(math.Round(
			(discountPercent / 100.0) * float64(resValue)),
		),
	}
}

// newTemperatureDiscount creates discount based on weather.
// Discount rules:
// - individual customers only
// - low outside temperature.
func newTemperatureDiscount(resValue int, customer bikerental.Customer, weather *bikerental.Weather) bikerental.Discount {
	if customer.Type != bikerental.CustomerTypeIndividual {
		return bikerental.Discount{}
	}
	if weather == nil || weather.Temperature >= 10 {
		return bikerental.Discount{}
	}

	return bikerental.Discount{
		Amount: int(math.Round(float64(resValue) * 0.05)),
	}
}

// newIncidentsDiscount creates discount based on incidents in the neighborhood.
// Discount rules:
// - individual customers only
// - incidents in neighborhood present.
func newIncidentsDiscount(resValue int, customer bikerental.Customer, incidents *bikerental.BikeIncidentsInfo) bikerental.Discount {
	if customer.Type != bikerental.CustomerTypeIndividual {
		return bikerental.Discount{}
	}

	if incidents == nil || incidents.NumberOfIncidents < 3 {
		return bikerental.Discount{}
	}

	discountPercent := 0.0
	if incidents.NumberOfIncidents >= 5 {
		discountPercent += 10.0
	} else {
		discountPercent += 5.0
	}
	return bikerental.Discount{
		Amount: int(math.Round(
			float64(resValue) * (discountPercent / 100.0),
		)),
	}
}

// selectOptimalDiscount chooses one discount that should be applied.
// Rules:
// - select discount with the greatest value.
func selectOptimalDiscount(discounts ...bikerental.Discount) bikerental.Discount {
	maxAmount := -math.MaxInt64
	var result bikerental.Discount
	for _, d := range discounts {
		if d.Amount > maxAmount {
			result = d
			maxAmount = d.Amount
		}
	}
	return result
}
