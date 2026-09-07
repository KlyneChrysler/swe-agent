package fixture

import "testing"

func TestExportCSVJoinsRows(t *testing.T) {
	report := Report{Rows: [][]string{{"a", "b"}, {"c", "d"}}}

	exported, err := report.Export(CSV)

	if err != nil {
		t.Fatalf("Export returned error: %v", err)
	}
	if exported != "a,b\nc,d" {
		t.Errorf("Export(CSV) = %q, want %q", exported, "a,b\nc,d")
	}
}

func TestExportRejectsUnknownFormat(t *testing.T) {
	_, err := Report{}.Export("xml")

	if err == nil {
		t.Fatal("Export accepted an unknown format")
	}
}

func TestFileNameUsesFormatExtension(t *testing.T) {
	name := Report{}.FileName("sales", JSON)

	if name != "sales.json" {
		t.Errorf("FileName = %q, want %q", name, "sales.json")
	}
}
