package fixture

import "testing"

func TestPriorityFor(t *testing.T) {
	cases := map[int]Priority{
		0:   Urgent,
		4:   Urgent,
		5:   Normal,
		48:  Normal,
		49:  Low,
		200: Low,
	}
	for hours, want := range cases {
		if got := PriorityFor(hours); got != want {
			t.Errorf("PriorityFor(%d) = %v, want %v", hours, got, want)
		}
	}
}
