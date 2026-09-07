package fixture

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Rate is the price of one unit of a currency in cents of the base currency.
type Rate struct {
	Currency string
	Cents    int64
}

type RateClient struct {
	http    *http.Client
	baseURL string
}

func NewRateClient(client *http.Client, baseURL string) RateClient {
	return RateClient{http: client, baseURL: baseURL}
}

func (c RateClient) Rate(ctx context.Context, currency string) (Rate, error) {
	response, err := c.get(ctx, "/rates/"+currency)
	if err != nil {
		return Rate{}, err
	}
	defer response.Body.Close()
	return decodeRate(response.Body)
}

func (c RateClient) get(ctx context.Context, path string) (*http.Response, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+path, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("build rate request %s: %w", path, err)
	}
	response, err := c.http.Do(request)
	if err != nil {
		return nil, fmt.Errorf("fetch rate %s: %w", path, err)
	}
	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		return nil, fmt.Errorf("fetch rate %s: unexpected status %d", path, response.StatusCode)
	}
	return response, nil
}

func decodeRate(body io.Reader) (Rate, error) {
	var rate Rate
	if err := json.NewDecoder(body).Decode(&rate); err != nil {
		return Rate{}, fmt.Errorf("decode rate: %w", err)
	}
	return rate, nil
}
