package fixture

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRateDecodesCentsFromUpstream(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"Currency":"EUR","Cents":108}`))
	}))
	defer upstream.Close()
	client := NewRateClient(upstream.Client(), upstream.URL)

	rate, err := client.Rate(context.Background(), "EUR")

	if err != nil {
		t.Fatalf("Rate returned error: %v", err)
	}
	if rate.Cents != 108 {
		t.Errorf("Rate cents = %d, want 108", rate.Cents)
	}
}

func TestRateReportsUnexpectedStatus(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
	}))
	defer upstream.Close()
	client := NewRateClient(upstream.Client(), upstream.URL)

	_, err := client.Rate(context.Background(), "EUR")

	if err == nil {
		t.Fatal("Rate accepted a 502 response")
	}
}
