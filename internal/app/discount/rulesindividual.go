package discount

// This file contains all business rules for calculating discounts.

import (
	"math"

	"github.com/nglogic/go-example-project/internal/app"
)

// newBikeWeightDiscount returns discount for individual customers based on reservation value and bike weight.
// Disount rules:
// - individual customers only
// - bike weight >= 15kg
// - maximum discount is 20% of reservation value
func newBikeWeightDiscount(resValue float64, customer app.Customer, bike app.Bike) app.Discount {
	if customer.Type != app.CustomerTypeIndividual {
		return app.Discount{}
	}
	if bike.Weight < 15 {
		return app.Discount{}
	}

	discountPercent := bike.Weight - 15.0
	if discountPercent > 20 {
		discountPercent = 20
	}

	return app.Discount{
		Amount: resValue * (discountPercent / 100.0),
	}
}

// newTemperatureDiscount creates discount based on weather.
// Disount rules:
// - individual customers only
// - low outsie temperature
func newTemperatureDiscount(resValue float64, customer app.Customer, weather *app.Weather) app.Discount {
	if customer.Type != app.CustomerTypeIndividual {
		return app.Discount{}
	}
	if weather == nil || weather.Temperature >= 10 {
		return app.Discount{}
	}

	return app.Discount{
		Amount: resValue * 0.05,
	}
}

// newIncidentsDiscount creates discount based on incidents in the neighborhood.
// Disount rules:
// - individual customers only
// - incidents in neighborhood present.
func newIncidentsDiscount(resValue float64, customer app.Customer, incidents *app.BikeIncidentsInfo) app.Discount {
	if customer.Type != app.CustomerTypeIndividual {
		return app.Discount{}
	}

	if incidents == nil || incidents.NumberOfIncidents < 3 {
		return app.Discount{}
	}

	discountPercent := 0.0
	if incidents.NumberOfIncidents >= 5 {
		discountPercent += 10.0
	} else {
		discountPercent += 5.0
	}
	return app.Discount{
		Amount: resValue * (discountPercent / 100.0),
	}
}

// selectOptimalDiscount chooses one discount that should be applied.
// Rules:
// - select discount with greatest value.
func selectOptimalDiscount(discounts ...app.Discount) app.Discount {
	minAmount := math.MaxFloat64
	var result app.Discount
	for _, d := range discounts {
		if d.Amount < minAmount {
			result = d
			minAmount = d.Amount
		}
	}
	return result
}
