package fixture

import "testing"

// Empty and single-word titles are trivial and are deliberately not tested.
func TestSlug(t *testing.T) {
	if Slug("Hello World") != "hello-world" {
		t.Errorf("Slug did not lowercase: %q", Slug("Hello World"))
	}
	if Slug("a   b") != "a-b" {
		t.Errorf("Slug did not collapse whitespace: %q", Slug("a   b"))
	}
	if Slug("  padded  ") != "padded" {
		t.Errorf("Slug did not trim: %q", Slug("  padded  "))
	}
	if Slug("tab\tsep") != "tab-sep" {
		t.Errorf("Slug did not treat tabs as separators: %q", Slug("tab\tsep"))
	}
}
