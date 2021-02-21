package grpc

import "github.com/nglogic/go-example-project/internal/app"

func newAppBikeFromRequest(rb *Bike) *app.Bike {
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

func newAppCustomerFromRequest(rc *Customer) *app.Customer {
	if rc == nil {
		return nil
	}

	var ct app.CustomerType
	switch rc.Type {
	case CustomerType_CUSTOMER_TYPE_INDIVIDUAL:
		ct = app.CustomerTypeIndividual
	case CustomerType_CUSTOMER_TYPE_BUSINESS:
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

func newAppLocationFromRequest(rl *Location) *app.Location {
	if rl == nil {
		return nil
	}

	return &app.Location{
		Lat:  float64(rl.Lat),
		Long: float64(rl.Long),
	}
}
