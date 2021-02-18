package server

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
