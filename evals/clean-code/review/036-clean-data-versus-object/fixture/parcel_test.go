package fixture

import (
	"errors"
	"testing"
)

var cebu = Address{Street: "12 Main St", City: "Cebu", PostalCode: "6000"}

func TestValidateAcceptsLightParcel(t *testing.T) {
	if err := NewParcel(500, cebu).Validate(); err != nil {
		t.Fatalf("Validate: %v", err)
	}
}

func TestValidateRejectsHeavyParcel(t *testing.T) {
	err := NewParcel(maxWeightGrams+1, cebu).Validate()
	if !errors.Is(err, ErrTooHeavy) {
		t.Fatalf("Validate = %v, want ErrTooHeavy", err)
	}
}

func TestIsBoundForMatchesCity(t *testing.T) {
	parcel := NewParcel(500, cebu)
	if !parcel.IsBoundFor("Cebu") || parcel.IsBoundFor("Manila") {
		t.Fatal("IsBoundFor did not match the destination city")
	}
}
