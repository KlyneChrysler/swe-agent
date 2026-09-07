package fixture

import "strings"

const unitedStates = "US"

type NormalizeAddress struct {
	Country string
}

func (n NormalizeAddress) Apply(line string) string {
	trimmed := strings.TrimSpace(line)
	if n.Country == unitedStates {
		return strings.ToUpper(trimmed)
	}
	return trimmed
}
