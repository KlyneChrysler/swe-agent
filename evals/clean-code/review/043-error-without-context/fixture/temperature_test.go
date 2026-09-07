package fixture

import "testing"

func TestParseCelsiusStripsSuffix(t *testing.T) {
	if got, err := ParseCelsius("21.5C"); err != nil || got != 21.5 {
		t.Fatalf("ParseCelsius = %v, %v; want 21.5", got, err)
	}
}

func TestParseCelsiusRejectsBadInput(t *testing.T) {
	for _, text := range []string{"warm", "-300C"} {
		if _, err := ParseCelsius(text); err == nil {
			t.Errorf("ParseCelsius(%q) accepted", text)
		}
	}
}
