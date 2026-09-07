package fixture

import "testing"

func TestIsSKU(t *testing.T) {
	cases := map[string]bool{
		"EL-00042": true,
		"el-00042": false,
		"EL-0042":  false,
		"":         false,
	}
	for text, want := range cases {
		if got := IsSKU(text); got != want {
			t.Errorf("IsSKU(%q) = %v, want %v", text, got, want)
		}
	}
}

func TestCategoryOfWellFormedSKU(t *testing.T) {
	if got := Category("EL-00042"); got != "EL" {
		t.Fatalf("Category = %q, want EL", got)
	}
}

func TestCategoryOfIllFormedSKUIsEmpty(t *testing.T) {
	if got := Category("nope"); got != "" {
		t.Fatalf("Category = %q, want empty", got)
	}
}
