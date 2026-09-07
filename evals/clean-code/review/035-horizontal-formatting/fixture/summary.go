package fixture

import "fmt"

type Summary struct {
	Average     int
	Highest     int
	Description string
}

func Summarize(readings []int) Summary {
	if len(readings) == 0 { return Summary{} }
	total   := 0
	highest := 0
	for _, reading := range readings { total += reading; if reading > highest { highest = reading } }
	average := total / len(readings)
	return Summary{Average: average, Highest: highest, Description: fmt.Sprintf("%d readings averaging %d with a peak of %d", len(readings), average, highest)}
}
