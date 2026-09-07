package fixture

import "testing"

func TestOvertime(t *testing.T) {
	cases := map[int]int{38: 0, 40: 0, 45: 5}
	for worked, want := range cases {
		if got := (Timesheet{HoursWorked: worked}).Overtime(); got != want {
			t.Errorf("Overtime(%d) = %d, want %d", worked, got, want)
		}
	}
}
