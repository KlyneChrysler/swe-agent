package fixture

import "testing"

func TestRenderHeaderCombinesTitleAndRegion(t *testing.T) {
	settings, err := LoadSettings([]byte(`{"title":"Sales","region":"EU","owner":"Ada"}`))
	if err != nil {
		t.Fatalf("LoadSettings returned error: %v", err)
	}

	header := RenderHeader(settings)

	if header != "Sales (EU)" {
		t.Errorf("RenderHeader = %q, want %q", header, "Sales (EU)")
	}
}

func TestRenderFooterNamesOwner(t *testing.T) {
	settings := map[string]any{"owner": "Ada"}

	footer := RenderFooter(settings)

	if footer != "Generated for Ada" {
		t.Errorf("RenderFooter = %q, want %q", footer, "Generated for Ada")
	}
}
