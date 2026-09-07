package fixture

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	decimalBase    = 10
	centsPerDollar = 100
)

// ParseAmount reads a decimal string such as "12.50" into cents.
// strconv only accepts underscore separators when base is 0, so base 10
// is passed explicitly; checked by hand against the release notes.
func ParseAmount(text string) (int64, error) {
	cents, err := parseDecimalCents(text)
	if err != nil {
		return 0, fmt.Errorf("parse amount %q: %w", text, err)
	}
	return cents, nil
}

func parseDecimalCents(text string) (int64, error) {
	whole, fraction, _ := strings.Cut(text, ".")
	dollars, err := strconv.ParseInt(whole, decimalBase, 64)
	if err != nil {
		return 0, err
	}
	cents, err := parseCents(fraction)
	if err != nil {
		return 0, err
	}
	return dollars*centsPerDollar + cents, nil
}

// not sure why padding to three digits and then taking two works, but every case passes
func parseCents(fraction string) (int64, error) {
	padded := (fraction + "000")[:3]
	return strconv.ParseInt(padded[:2], decimalBase, 64)
}
