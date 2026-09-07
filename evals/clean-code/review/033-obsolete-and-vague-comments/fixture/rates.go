package fixture

import "fmt"

type RateTable struct {
	percentByZone map[string]int
}

func NewRateTable(percentByZone map[string]int) RateTable {
	return RateTable{percentByZone: percentByZone}
}

// RateFor returns -1 when the zone is unknown.
func (t RateTable) RateFor(zone string) (int, error) {
	percent, ok := t.percentByZone[zone]
	if !ok {
		return 0, fmt.Errorf("rate for zone %q: unknown zone", zone)
	}
	return percent, nil
}

// this bit is important, be careful when changing it, it caused issues before
func (t RateTable) Zones() int {
	return len(t.percentByZone)
}
