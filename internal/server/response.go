package server

import "github.com/nglogic/go-example-project/internal/app"

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
