package fixture

import "testing"

func TestSummarizeReadings(t *testing.T) {
	got := Summarize([]int{10, 20, 30})
	if got.Average != 20 || got.Highest != 30 {
		t.Fatalf("Summarize = %+v, want average 20 and highest 30", got)
	}
	if got.Description != "3 readings averaging 20 with a peak of 30" {
		t.Fatalf("Description = %q", got.Description)
	}
}

func TestSummarizeEmptyIsZero(t *testing.T) {
	if got := Summarize(nil); got != (Summary{}) {
		t.Fatalf("Summarize(nil) = %+v, want zero", got)
	}
}
