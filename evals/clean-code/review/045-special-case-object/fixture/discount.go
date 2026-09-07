package fixture

import "errors"

const percentDenominator = 100

var ErrNoDiscount = errors.New("discount for sku: none configured")

type Discount struct {
	Percent int
}

type Catalog struct {
	discounts map[string]Discount
}

func NewCatalog(discounts map[string]Discount) Catalog {
	return Catalog{discounts: discounts}
}

func (c Catalog) DiscountFor(sku string) (Discount, error) {
	discount, ok := c.discounts[sku]
	if !ok {
		return Discount{}, ErrNoDiscount
	}
	return discount, nil
}

func (c Catalog) PriceCents(sku string, baseCents int) int {
	discount, err := c.DiscountFor(sku)
	if errors.Is(err, ErrNoDiscount) {
		return baseCents
	}
	return baseCents - baseCents*discount.Percent/percentDenominator
}
