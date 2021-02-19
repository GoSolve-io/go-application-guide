package server

import (
	"github.com/nglogic/go-example-project/internal/app"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func newListBikesResponse(bikes []app.Bike) *ListBikesResponse {
	respBikes := make([]*Bike, 0, len(bikes))
	for _, b := range bikes {
		respBikes = append(respBikes, newResponseBike(&b))
	}

	return &ListBikesResponse{
		Bikes: respBikes,
	}
}

func newResponseBike(b *app.Bike) *Bike {
	if b == nil {
		return nil
	}
	return &Bike{
		Id:           b.ID,
		ModelName:    b.ModelName,
		Weight:       float32(b.Weight),
		PricePerHour: float32(b.PricePerHour),
	}
}

func newCreateReservationResponse(r *app.ReservationResponse) *CreateReservationResponse {
	if r == nil {
		return nil
	}

	var status ReservationStatus
	switch r.Status {
	case app.ReservationStatusApproved:
		status = ReservationStatus_RESERVATION_STATUS_APPROVED
	case app.ReservationStatusRejected:
		status = ReservationStatus_RESERVATION_STATUS_REJECTED
	default:
		status = ReservationStatus_RESERVATION_STATUS_UNKNOWN
	}

	return &CreateReservationResponse{
		Reservation: newResponseReservation(r.Reservation),
		Status:      status,
		Reason:      r.Reason,
	}
}

func newResponseReservation(r *app.Reservation) *Reservation {
	if r == nil {
		return nil
	}
	return &Reservation{
		Id:              r.ID,
		Customer:        newResponseCustomer(&r.Customer),
		Bike:            newResponseBike(&r.Bike),
		StartTime:       timestamppb.New(r.StartTime),
		EndTime:         timestamppb.New(r.EndTime),
		TotalValue:      float32(r.TotalValue),
		AppliedDiscount: float32(r.AppliedDiscount),
	}
}

func newResponseCustomer(c *app.Customer) *Customer {
	if c == nil {
		return nil
	}
	var t CustomerType
	switch c.Type {
	case app.CustomerTypeIndividual:
		t = CustomerType_CUSTOMER_TYPE_INDIVIDUAL
	case app.CustomerTypeBuisiness:
		t = CustomerType_CUSTOMER_TYPE_BUSINESS
	default:
		t = CustomerType_CUSTOMER_TYPE_UNKNOWN
	}
	return &Customer{
		Id:        c.ID,
		Type:      t,
		FirstName: c.FirstName,
		Surname:   c.Surname,
		Email:     c.Email,
	}
}
