package fixture

import "testing"

func TestParseAmountReadsDollarsAndCents(t *testing.T) {
	cases := []struct {
		text string
		want int64
	}{
		{"12.50", 1250},
		{"7", 700},
		{"0.5", 50},
	}
	for _, c := range cases {
		got, err := ParseAmount(c.text)
		if err != nil {
			t.Fatalf("ParseAmount(%q) returned error: %v", c.text, err)
		}
		if got != c.want {
			t.Errorf("ParseAmount(%q) = %d, want %d", c.text, got, c.want)
		}
	}
}

func TestParseAmountRejectsLetters(t *testing.T) {
	_, err := ParseAmount("12.x")

	if err == nil {
		t.Fatal("ParseAmount accepted a non-numeric fraction")
	}
}
