package fixture

import "regexp"

const categoryLength = 2

// skuPattern matches a warehouse SKU: a two-letter category, a dash, and a
// five-digit sequence number, for example "EL-00042".
var skuPattern = regexp.MustCompile(`^[A-Z]{2}-[0-9]{5}$`)

// IsSKU reports whether text is a well-formed warehouse SKU. It does not
// check that the SKU exists in the catalog.
func IsSKU(text string) bool {
	return skuPattern.MatchString(text)
}

// Category returns the two-letter category prefix of sku, or "" when sku is
// not well-formed.
func Category(sku string) string {
	if !IsSKU(sku) {
		return ""
	}
	return sku[:categoryLength]
}
