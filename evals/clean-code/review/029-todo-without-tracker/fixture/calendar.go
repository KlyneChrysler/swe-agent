package fixture

const daysPerYear = 365

// TODO: leap years are wrong here, fix at some point
func DaysBetweenYears(from, to int) int {
	return (to - from) * daysPerYear
}
