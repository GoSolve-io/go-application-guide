package discount

// This file contains all business rules for calculating discounts.

import (
	"math"

	"github.com/nglogic/go-application-guide/internal/app/bikerental"
)

// newBusinessCustomerDiscount returns discounts for business customers.
// Discount rules:
// - business customers only
// - minimum reservation value: 100
// - discount value: 5% of reservation value
func newBusinessCustomerDiscount(resValue int, customer bikerental.Customer) bikerental.Discount {
	if customer.Type != bikerental.CustomerTypeBuisiness {
		return bikerental.Discount{}
	}
	if resValue < 100 {
		return bikerental.Discount{}
	}
	return bikerental.Discount{
		Amount: int(math.Round(
			0.05 * float64(resValue),
		)),
	}
}
