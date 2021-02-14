package app

import "context"

// Weather represents weather data.
type Weather struct {
	Temperature float64
}

// WeatherRequest is a request for weather data in a location.
type WeatherRequest struct {
	Location Location
}

// WeatherService provides weather data.
type WeatherService interface {
	GetWeather(context.Context, WeatherRequest) (*Weather, error)
}
