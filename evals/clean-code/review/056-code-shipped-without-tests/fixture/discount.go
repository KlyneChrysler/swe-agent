package fixture

// Tier is a customer loyalty level; higher tiers earn larger discounts.
type Tier int

const (
	Standard Tier = iota
	Silver
	Gold
)

const (
	silverPercent = 5
	goldPercent   = 10
	percent       = 100
)

var tierNames = []string{"standard", "silver", "gold"}

func (t Tier) String() string {
	return tierNames[t]
}

func Discount(tier Tier, cents int64) int64 {
	return cents * discountPercent(tier) / percent
}

func discountPercent(tier Tier) int64 {
	switch tier {
	case Silver:
		return silverPercent
	case Gold:
		return goldPercent
	}
	return 0
}
