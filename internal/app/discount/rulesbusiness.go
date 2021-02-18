package discount

// This file contains all business rules for calculating discounts.

import (
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
