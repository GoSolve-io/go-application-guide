package discounts

// This file contains all business rules for calculating discounts.

import (
	"math"

	"github.com/nglogic/go-example-project/internal/app"
)

// newBusinessCustomerDiscount returns discounts for business customers.
// Discount rules:
// - business customers only
// - minimum reservation value: 100
// - discount value: 5% of reservation value
func newBusinessCustomerDiscount(resValue float64, customer app.Customer) app.Discount {
	if customer.Type != app.CustomerTypeBuisiness {
		return app.Discount{}
	}
	if resValue < 100 {
		return app.Discount{}
	}
	return app.Discount{Amount: 0.05 * resValue}
}

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

// newEnvironmentalDiscount creates discount based on weather and incidents in neighbourhood.
// Disount rules:
// - individual customers only
// - low outsie temperature
func newEnvironmentalDiscount(resValue float64, customer app.Customer, weather app.Weather, incidents app.BikeIncidentsInfo) app.Discount {
	if customer.Type != app.CustomerTypeIndividual {
		return app.Discount{}
	}

	discountPercent := 0.0
	if weather.Temperature < 10 {
		discountPercent += 5.0
	}
	if incidents.NumberOfIncidents >= 5 {
		discountPercent += 10.0
	} else if incidents.NumberOfIncidents >= 3 {
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
