package fixture

import "strings"

// Slug turns a title into a lowercase, hyphen-separated URL segment.
func Slug(title string) string {
	words := strings.Fields(strings.ToLower(title))
	return strings.Join(words, "-")
}
