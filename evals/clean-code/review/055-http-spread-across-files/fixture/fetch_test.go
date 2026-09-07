package fixture

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func serve(t *testing.T, status int, body string) *httptest.Server {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write([]byte(body))
	}))
	t.Cleanup(server.Close)
	return server
}

func TestFetchPriceDecodesCents(t *testing.T) {
	upstream := serve(t, http.StatusOK, `{"SKU":"A1","Cents":499}`)

	price, err := FetchPrice(upstream.URL, "A1")

	if err != nil {
		t.Fatalf("FetchPrice returned error: %v", err)
	}
	if price.Cents != 499 {
		t.Errorf("FetchPrice cents = %d, want 499", price.Cents)
	}
}

func TestFetchStockParsesCount(t *testing.T) {
	upstream := serve(t, http.StatusOK, "12\n")

	count, err := FetchStock(Warehouse{URL: upstream.URL, Token: "t"}, "A1")

	if err != nil {
		t.Fatalf("FetchStock returned error: %v", err)
	}
	if count != 12 {
		t.Errorf("FetchStock = %d, want 12", count)
	}
}

func TestIsUpstreamHealthyRequiresOK(t *testing.T) {
	upstream := serve(t, http.StatusServiceUnavailable, "")

	if IsUpstreamHealthy(upstream.URL) {
		t.Error("IsUpstreamHealthy reported a 503 upstream as healthy")
	}
}
