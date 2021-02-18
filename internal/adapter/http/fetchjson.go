package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Doer is a utility interface. It's implemented by http.Client.
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

// GetJSON fetches json data using HTTP GET request.
// Json from response is unmarshalled to the `result` object (it usually should be a pointer!).
func GetJSON(
	ctx context.Context,
	doer Doer,
	timeout time.Duration,
	url string,
	result interface{},
) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("couldn't create http request: %w", err)
	}

	resp, err := doer.Do(req)
	if err != nil {
		return fmt.Errorf("http get '%s': %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid http status: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("decoding json response from metaweather: %w", err)
	}

	return nil
}
