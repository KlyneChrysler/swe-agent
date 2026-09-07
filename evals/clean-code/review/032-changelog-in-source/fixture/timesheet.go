package fixture

// Author: J. Alvarez
// Created: 2024-02-11
// Modified: 2024-03-02 - added Overtime
// Modified: 2024-05-19 - fixed rounding in Overtime

const standardWeekHours = 40

type Timesheet struct {
	HoursWorked int
}

func (t Timesheet) Overtime() int {
	if t.HoursWorked <= standardWeekHours {
		return 0
	}
	return t.HoursWorked - standardWeekHours
}
