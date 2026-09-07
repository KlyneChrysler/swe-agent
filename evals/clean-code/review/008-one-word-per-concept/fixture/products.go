package fixture

import "strings"

type Product struct {
	SKU   string
	Cents int
}

type Catalog struct {
	bySKU map[string]Product
}

func (c Catalog) GetProduct(sku string) (Product, bool) {
	product, ok := c.bySKU[strings.ToUpper(sku)]
	return product, ok
}
