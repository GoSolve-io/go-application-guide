package grpc

import (
	"github.com/nglogic/go-application-guide/internal/app/bikerental"
	"github.com/nglogic/go-application-guide/pkg/api/bikerentalv1"
)

func newAppBikeFromRequestData(data *bikerentalv1.BikeData) *bikerental.Bike {
	return &bikerental.Bike{
		ModelName:    data.ModelName,
		Weight:       float64(data.Weight),
		PricePerHour: int(data.PricePerHour),
	}
}

func newAppCustomerFromRequest(rc *bikerentalv1.Customer) *bikerental.Customer {
	if rc == nil {
		return nil
	}

	data := rc.GetData()
	ct := bikerental.CustomerTypeUnknown
	switch data.GetType() {
	case bikerentalv1.CustomerType_CUSTOMER_TYPE_INDIVIDUAL:
		ct = bikerental.CustomerTypeIndividual
	case bikerentalv1.CustomerType_CUSTOMER_TYPE_BUSINESS:
		ct = bikerental.CustomerTypeBusiness
	}

	return &bikerental.Customer{
		ID:        rc.Id,
		Type:      ct,
		FirstName: data.GetFirstName(),
		Surname:   data.GetSurname(),
		Email:     data.GetEmail(),
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
