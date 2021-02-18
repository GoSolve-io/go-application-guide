package incidents

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/nglogic/go-example-project/internal/adapter/http"
	"github.com/nglogic/go-example-project/internal/app"
)

const (
	// Maximum number of incidents that GetIncidents can report
	// This is used to cut response size from bikewise api, because they don't return total number
	// and we have to count returned objects.
	maxIncidents = 50
)

// Adapter uses bikewise.org API for bike incidents data.
// See: https://www.bikewise.org/documentation/api_v2#!/incidents/GET_version_incidents_format_get_0
type Adapter struct {
	address  string
	timeout  time.Duration
	httpDoer http.Doer
}

// NewAdapter creates new adapter instance.
func NewAdapter(address string, timeout time.Duration, httpDoer http.Doer) (*Adapter, error) {
	if address == "" {
		return nil, errors.New("address is required")
	}
	if timeout == 0 {
		return nil, errors.New("timeout is required")
	}
	if httpDoer == nil {
		return nil, errors.New("http doer is required")
	}

	return &Adapter{
		address:  address,
		timeout:  timeout,
		httpDoer: httpDoer,
	}, nil
}

// GetIncidents return number of bike incidents in a location.
// Maximum returned value will be `maxIncidents`.
func (a *Adapter) GetIncidents(ctx context.Context, req app.BikeIncidentsRequest) (*app.BikeIncidentsInfo, error) {
	urlVal := fmt.Sprintf("%s/v2/locations", a.address)
	query := url.Values{
		"proximity": []string{
			fmt.Sprintf("%f,%f", req.Location.Lat, req.Location.Long),
		},
		"proximity_square": []string{
			fmt.Sprintf("%f", req.Proximity),
		},
		"limit": []string{strconv.Itoa(maxIncidents)},
	}

	var resp bikewiseLocationsResponse
	if err := http.GetJSON(
		ctx,
		a.httpDoer,
		a.timeout,
		fmt.Sprintf("%s?%s", urlVal, query.Encode()),
		&resp,
	); err != nil {
		return nil, fmt.Errorf("fetching data from bikewise: %w", err)
	}

	return &app.BikeIncidentsInfo{
		Location:          req.Location,
		Proximity:         req.Proximity,
		NumberOfIncidents: len(resp.Features),
	}, nil
}

type bikewiseLocationsResponse struct {
	// Features is just a list of some object. We only care about count, so internal structure is irrelevant.
	Features []json.RawMessage `json:"features"`
}
