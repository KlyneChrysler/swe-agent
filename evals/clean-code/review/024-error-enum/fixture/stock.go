package fixture

import "strconv"

type Status int

const (
	OK Status = iota
	NotFound
	Malformed
)

type Stock struct {
	onHand map[string]int
}

func NewStock(onHand map[string]int) *Stock {
	return &Stock{onHand: onHand}
}

func (s *Stock) Restock(sku, quantityText string) Status {
	quantity, err := strconv.Atoi(quantityText)
	if err != nil {
		return Malformed
	}
	if _, known := s.onHand[sku]; !known {
		return NotFound
	}
	s.onHand[sku] += quantity
	return OK
}

func Describe(status Status) string {
	switch status {
	case NotFound:
		return "unknown sku"
	case Malformed:
		return "quantity is not a number"
	default:
		return "ok"
	}
}
