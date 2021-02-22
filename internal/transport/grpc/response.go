package grpc

import (
	"github.com/nglogic/go-example-project/internal/app"
	v1 "github.com/nglogic/go-example-project/pkg/api/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func newListBikesResponse(bikes []app.Bike) *v1.ListBikesResponse {
	respBikes := make([]*v1.Bike, 0, len(bikes))
	for _, b := range bikes {
		respBikes = append(respBikes, newResponseBike(&b))
	}

	return &v1.ListBikesResponse{
		Bikes: respBikes,
	}
}

func newResponseBike(b *app.Bike) *v1.Bike {
	if b == nil {
		return nil
	}
	return &v1.Bike{
		Id:           b.ID,
		ModelName:    b.ModelName,
		Weight:       float32(b.Weight),
		PricePerHour: float32(b.PricePerHour),
	}
}

func newCreateReservationResponse(r *app.ReservationResponse) *v1.CreateReservationResponse {
	if r == nil {
		return nil
	}

	return &v1.CreateReservationResponse{
		Reservation: newResponseReservation(r.Reservation),
		Status:      newResponseReservationStatus(r.Status),
		Reason:      r.Reason,
	}
}

func newResponseReservation(r *app.Reservation) *v1.Reservation {
	if r == nil {
		return nil
	}
	return &v1.Reservation{
		Id:              r.ID,
		Status:          newResponseReservationStatus(r.Status),
		Customer:        newResponseCustomer(&r.Customer),
		Bike:            newResponseBike(&r.Bike),
		StartTime:       timestamppb.New(r.StartTime),
		EndTime:         timestamppb.New(r.EndTime),
		TotalValue:      float32(r.TotalValue),
		AppliedDiscount: float32(r.AppliedDiscount),
	}
}

func newResponseCustomer(c *app.Customer) *v1.Customer {
	if c == nil {
		return nil
	}
	var t v1.CustomerType
	switch c.Type {
	case app.CustomerTypeIndividual:
		t = v1.CustomerType_CUSTOMER_TYPE_INDIVIDUAL
	case app.CustomerTypeBuisiness:
		t = v1.CustomerType_CUSTOMER_TYPE_BUSINESS
	default:
		t = v1.CustomerType_CUSTOMER_TYPE_UNKNOWN
	}
	return &v1.Customer{
		Id:        c.ID,
		Type:      t,
		FirstName: c.FirstName,
		Surname:   c.Surname,
		Email:     c.Email,
	}
}

func newResponseReservationStatus(s app.ReservationStatus) v1.ReservationStatus {
	var status v1.ReservationStatus
	switch s {
	case app.ReservationStatusApproved:
		status = v1.ReservationStatus_RESERVATION_STATUS_APPROVED
	case app.ReservationStatusRejected:
		status = v1.ReservationStatus_RESERVATION_STATUS_REJECTED
	case app.ReservationStatusCanceled:
		status = v1.ReservationStatus_RESERVATION_STATUS_CANCELLED
	default:
		status = v1.ReservationStatus_RESERVATION_STATUS_UNKNOWN
	}
	return status
}
