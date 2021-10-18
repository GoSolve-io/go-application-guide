package weather

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	ahttp "github.com/nglogic/go-application-guide/internal/adapter/http"
	"github.com/nglogic/go-application-guide/internal/app/bikerental"
)

// Adapter uses metaweather service for providing weather data.
type Adapter struct {
	// address valid value can be "https://www.metaweather.com"
	address  string
	timeout  time.Duration
	httpDoer ahttp.Doer
}

// NewAdapter creates new adapter instance.
func NewAdapter(address string, timeout time.Duration, httpDoer ahttp.Doer) (*Adapter, error) {
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

// GetWeather fetches weather data for a location.
func (a *Adapter) GetWeather(ctx context.Context, req bikerental.WeatherRequest) (*bikerental.Weather, error) {
	locID, err := a.fetchLocationID(ctx, req.Location)
	if err != nil {
		return nil, fmt.Errorf("fetching location id by coordinates (%s): %w", req.Location.String(), err)
	}
	if locID == 0 {
		return nil, nil
	}

	weather, err := a.fetchCurrentWeather(ctx, locID)
	if err != nil {
		return nil, fmt.Errorf("fetching weather for location (id=%d): %w", locID, err)
	}
	if weather == nil {
		return nil, nil
	}
	return &bikerental.Weather{
		Temperature: weather.TheTemp,
	}, nil
}

func (a *Adapter) fetchLocationID(ctx context.Context, loc bikerental.Location) (int, error) {
	urlVal := fmt.Sprintf("%s/api/location/search/", a.address)
	query := url.Values{
		"lattlong": []string{
			fmt.Sprintf("%f,%f", loc.Lat, loc.Long),
		},
	}

	var result []locationEntry
	err := ahttp.GetJSON(
		ctx,
		a.httpDoer,
		a.timeout,
		fmt.Sprintf("%s?%s", urlVal, query.Encode()),
		&result,
	)
	if err != nil {
		return 0, fmt.Errorf("fetching data from metaweather: %w", err)
	}

	if len(result) == 0 {
		return 0, nil
	}
	return result[0].Woeid, nil
}

func (a *Adapter) fetchCurrentWeather(ctx context.Context, locID int) (*weatherEntry, error) {
	if locID == 0 {
		return nil, errors.New("got 0 location id")
	}

	urlVal := fmt.Sprintf("%s/api/location/%d/", a.address, locID)

	var resp weatherResponse
	if err := ahttp.GetJSON(ctx, a.httpDoer, a.timeout, urlVal, &resp); err != nil {
		return nil, fmt.Errorf("fetching data from metaweather: %w", err)
	}

	if len(resp.ConsolidatedWeather) == 0 {
		return nil, nil
	}
	return &resp.ConsolidatedWeather[0], nil
}

type locationEntry struct {
	Woeid int `json:"woeid"`
}

type weatherResponse struct {
	ConsolidatedWeather []weatherEntry `json:"consolidated_weather"`
}

type weatherEntry struct {
	AirPressure          float64   `json:"air_pressure"`
	ApplicableDate       string    `json:"applicable_date"`
	Created              time.Time `json:"created"`
	Humidity             int       `json:"humidity"`
	ID                   int       `json:"id"`
	MaxTemp              float64   `json:"max_temp"`
	MinTemp              float64   `json:"min_temp"`
	Predictability       int       `json:"predictability"`
	TheTemp              float64   `json:"the_temp"`
	Visibility           float64   `json:"visibility"`
	WeatherStateAbbr     string    `json:"weather_state_abbr"`
	WeatherStateName     string    `json:"weather_state_name"`
	WindDirection        float64   `json:"wind_direction"`
	WindDirectionCompass string    `json:"wind_direction_compass"`
	WindSpeed            float64   `json:"wind_speed"`
}
