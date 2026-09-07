package fixture

import (
	"errors"
	"fmt"
	"strconv"
)

const percentDenominator = 100

type Invoice struct {
	TotalCents int
}

func (self *Invoice) Apply_Discount(percent int) {
	self.TotalCents -= self.TotalCents * percent / percentDenominator
}

func (inv *Invoice) AddLine(cents int) {
	inv.TotalCents += cents
}

func ParseCents(text string) (int, error) {
	cents, err := strconv.Atoi(text)
	if err == nil {
		return cents, nil
	} else {
		return 0, errors.New("parse cents " + text + ": " + err.Error())
	}
}

func ParsePercent(text string) (int, error) {
	percent, err := strconv.Atoi(text)
	if err != nil {
		return 0, fmt.Errorf("parse percent %q: %w", text, err)
	}
	return percent, nil
}
