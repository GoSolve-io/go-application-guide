package grpc

import (
	"github.com/nglogic/go-example-project/internal/app/bikerental"
	"github.com/nglogic/go-example-project/pkg/api/bikerentalv1"
)

func newAppBikeFromRequest(rb *bikerentalv1.Bike) *bikerental.Bike {
	if rb == nil {
		return nil
	}
	return &bikerental.Bike{
		ID:           rb.Id,
		ModelName:    rb.ModelName,
		Weight:       float64(rb.Weight),
		PricePerHour: float64(rb.PricePerHour),
	}
}

func newAppCustomerFromRequest(rc *bikerentalv1.Customer) *bikerental.Customer {
	if rc == nil {
		return nil
	}

	var ct bikerental.CustomerType
	switch rc.Type {
	case bikerentalv1.CustomerType_CUSTOMER_TYPE_INDIVIDUAL:
		ct = bikerental.CustomerTypeIndividual
	case bikerentalv1.CustomerType_CUSTOMER_TYPE_BUSINESS:
		ct = bikerental.CustomerTypeBuisiness
	}

	return &bikerental.Customer{
		ID:        rc.Id,
		Type:      ct,
		FirstName: rc.FirstName,
		Surname:   rc.Surname,
		Email:     rc.Email,
	}
}

func newAppLocationFromRequest(rl *bikerentalv1.Location) *bikerental.Location {
	if rl == nil {
		return nil
	}

	return &bikerental.Location{
		Lat:  float64(rl.Lat),
		Long: float64(rl.Long),
	}
}
