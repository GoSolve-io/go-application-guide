package server

import "github.com/nglogic/go-example-project/internal/app"

func newListBikesResponse(bikes []app.Bike) *ListBikesResponse {
	respBikes := make([]*Bike, 0, len(bikes))
	for _, b := range bikes {
		rb := Bike{
			Id:           b.ID,
			ModelName:    b.ModelName,
			Weight:       float32(b.Weight),
			PricePerHour: float32(b.PricePerHour),
		}
		respBikes = append(respBikes, &rb)
	}

	return &ListBikesResponse{
		Bikes: respBikes,
	}
}
