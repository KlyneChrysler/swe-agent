package fixture

import (
	"errors"
	"strconv"
	"strings"
)

const (
	celsiusSuffix       = "C"
	absoluteZeroCelsius = -273.15
)

var errInvalid = errors.New("invalid")

func ParseCelsius(text string) (float64, error) {
	value, err := strconv.ParseFloat(strings.TrimSuffix(text, celsiusSuffix), 64)
	if err != nil {
		return 0, errors.New("error")
	}
	if value < absoluteZeroCelsius {
		return 0, errInvalid
	}
	return value, nil
}
