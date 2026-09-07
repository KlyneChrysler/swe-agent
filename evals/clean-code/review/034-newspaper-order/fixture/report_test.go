package fixture

import "testing"

func TestRenderPadsNamesAndFormatsCents(t *testing.T) {
	got := Render([]Line{{Name: "seat", Cents: 1205}, {Name: "tax", Cents: 50}})
	want := "seat                    $12.05\ntax                     $0.50"
	if got != want {
		t.Fatalf("Render = %q, want %q", got, want)
	}
}
