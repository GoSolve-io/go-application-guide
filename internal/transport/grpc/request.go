package grpc

import (
	"github.com/nglogic/go-example-project/internal/app"
	"github.com/nglogic/go-example-project/pkg/api/bikerentalv1"
)

func newAppBikeFromRequest(rb *bikerentalv1.Bike) *app.Bike {
	if rb == nil {
		return nil
	}
	return &app.Bike{
		ID:           rb.Id,
		ModelName:    rb.ModelName,
		Weight:       float64(rb.Weight),
		PricePerHour: float64(rb.PricePerHour),
	}
}

func newAppCustomerFromRequest(rc *bikerentalv1.Customer) *app.Customer {
	if rc == nil {
		return nil
	}

	var ct app.CustomerType
	switch rc.Type {
	case bikerentalv1.CustomerType_CUSTOMER_TYPE_INDIVIDUAL:
		ct = app.CustomerTypeIndividual
	case bikerentalv1.CustomerType_CUSTOMER_TYPE_BUSINESS:
		ct = app.CustomerTypeBuisiness
	}

	return &app.Customer{
		ID:        rc.Id,
		Type:      ct,
		FirstName: rc.FirstName,
		Surname:   rc.Surname,
		Email:     rc.Email,
	}
}

func newAppLocationFromRequest(rl *bikerentalv1.Location) *app.Location {
	if rl == nil {
		return nil
	}

	return &app.Location{
		Lat:  float64(rl.Lat),
		Long: float64(rl.Long),
	}
}
