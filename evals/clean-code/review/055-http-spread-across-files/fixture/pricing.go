package fixture

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Price is a catalogue price in cents.
type Price struct {
	SKU   string
	Cents int64
}

func FetchPrice(catalogueURL, sku string) (Price, error) {
	response, err := http.Get(catalogueURL + "/prices/" + sku)
	if err != nil {
		return Price{}, fmt.Errorf("fetch price %s: %w", sku, err)
	}
	defer response.Body.Close()
	return decodePrice(response.Body)
}

func decodePrice(body io.Reader) (Price, error) {
	var price Price
	if err := json.NewDecoder(body).Decode(&price); err != nil {
		return Price{}, fmt.Errorf("decode price: %w", err)
	}
	return price, nil
}
