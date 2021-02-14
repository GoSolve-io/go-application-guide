package app

import "context"

// BikeIncidentsInfo represents information about bike incidents within a square of size `proximity`, centered at `location`.
type BikeIncidentsInfo struct {
	Location          Location
	Proximity         float64
	NumberOfIncidents int
}

// BikeIncidentsRequest is a request for incidents data.
type BikeIncidentsRequest struct {
	// Location defines center point of square in which we search for incidents.
	Location Location
	// Proximity defines square size in km, in which we search for incidents.
	Proximity float64
}

// BikeIncidentsService provides information about bike incidents.
type BikeIncidentsService interface {
	GetIncidents(context.Context, BikeIncidentsRequest) (*BikeIncidentsInfo, error)
}
