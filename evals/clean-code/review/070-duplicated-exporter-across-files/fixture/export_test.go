package fixture

import "testing"

func TestExportCSVWritesHeaderThenRows(t *testing.T) {
	exported := ExportCSV([]string{"id", "name"}, [][]string{{"1", "Ada"}})

	if exported != "id,name\n1,Ada\n" {
		t.Errorf("ExportCSV = %q", exported)
	}
}

func TestExportTSVWritesHeaderThenRows(t *testing.T) {
	exported := ExportTSV([]string{"id", "name"}, [][]string{{"1", "Ada"}})

	if exported != "id\tname\n1\tAda\n" {
		t.Errorf("ExportTSV = %q", exported)
	}
}
