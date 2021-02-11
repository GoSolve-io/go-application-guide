package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/nglogic/go-example-project/internal/app"
)

// Adapter uses metaweather service for providing weather data.
type Adapter struct {
	// address valid value can be "https://www.metaweather.com"
	address    string
	timeout    time.Duration
	httpClient *http.Client
}

// GetWeather fetches weather data for a location.
func (a *Adapter) GetWeather(ctx context.Context, req app.WeatherRequest) (*app.Weather, error) {
	response, err := a.fetch(ctx, req.Location)
	if err != nil {
		return nil, err
	}
	if len(response) == 0 {
		return nil, nil
	}
	return &app.Weather{
		// Temperature: response[0]., TODO
	}, nil
}

func (a *Adapter) fetch(ctx context.Context, loc app.Location) (metaweatherLocationResponse, error) {
	urlVal := fmt.Sprintf("%s/api/location/search/", a.address)
	query := url.Values{
		"lattlong": []string{
			fmt.Sprintf("%f", loc.Lat),
			fmt.Sprintf("%f", loc.Lng),
		},
	}

	ctx, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	resp, err := a.httpClient.Get(fmt.Sprintf("%s?%s", urlVal, query.Encode()))
	if err != nil {
		return metaweatherLocationResponse{}, fmt.Errorf("fetching data from metaweather: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return metaweatherLocationResponse{}, fmt.Errorf("metaweather returned invalid http status: %d", resp.StatusCode)
	}

	var result metaweatherLocationResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return metaweatherLocationResponse{}, fmt.Errorf("decoding json response from metaweather: %w", err)
	}

	return result, nil
}

type metaweatherLocationResponse []metaweatherLocation

type metaweatherLocation struct {
	Woeid int `json:"woeid"`
}
