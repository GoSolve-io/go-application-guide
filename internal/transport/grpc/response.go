package grpc

import (
	"github.com/nglogic/go-example-project/internal/app"
	"github.com/nglogic/go-example-project/pkg/api/bikerentalv1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func newListBikesResponse(bikes []app.Bike) *bikerentalv1.ListBikesResponse {
	respBikes := make([]*bikerentalv1.Bike, 0, len(bikes))
	for _, b := range bikes {
		respBikes = append(respBikes, newResponseBike(&b))
	}

	return &bikerentalv1.ListBikesResponse{
		Bikes: respBikes,
	}
}

func newResponseBike(b *app.Bike) *bikerentalv1.Bike {
	if b == nil {
		return nil
	}
	return &bikerentalv1.Bike{
		Id:           b.ID,
		ModelName:    b.ModelName,
		Weight:       float32(b.Weight),
		PricePerHour: float32(b.PricePerHour),
	}
}

func newCreateReservationResponse(r *app.ReservationResponse) *bikerentalv1.CreateReservationResponse {
	if r == nil {
		return nil
	}

	return &bikerentalv1.CreateReservationResponse{
		Reservation: newResponseReservation(r.Reservation),
		Status:      newResponseReservationStatus(r.Status),
		Reason:      r.Reason,
	}
}

func newResponseReservation(r *app.Reservation) *bikerentalv1.Reservation {
	if r == nil {
		return nil
	}
	return &bikerentalv1.Reservation{
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

func newResponseCustomer(c *app.Customer) *bikerentalv1.Customer {
	if c == nil {
		return nil
	}
	var t bikerentalv1.CustomerType
	switch c.Type {
	case app.CustomerTypeIndividual:
		t = bikerentalv1.CustomerType_CUSTOMER_TYPE_INDIVIDUAL
	case app.CustomerTypeBuisiness:
		t = bikerentalv1.CustomerType_CUSTOMER_TYPE_BUSINESS
	default:
		t = bikerentalv1.CustomerType_CUSTOMER_TYPE_UNKNOWN
	}
	return &bikerentalv1.Customer{
		Id:        c.ID,
		Type:      t,
		FirstName: c.FirstName,
		Surname:   c.Surname,
		Email:     c.Email,
	}
}

func newResponseReservationStatus(s app.ReservationStatus) bikerentalv1.ReservationStatus {
	var status bikerentalv1.ReservationStatus
	switch s {
	case app.ReservationStatusApproved:
		status = bikerentalv1.ReservationStatus_RESERVATION_STATUS_APPROVED
	case app.ReservationStatusRejected:
		status = bikerentalv1.ReservationStatus_RESERVATION_STATUS_REJECTED
	case app.ReservationStatusCanceled:
		status = bikerentalv1.ReservationStatus_RESERVATION_STATUS_CANCELLED
	default:
		status = bikerentalv1.ReservationStatus_RESERVATION_STATUS_UNKNOWN
	}
	return status
}
