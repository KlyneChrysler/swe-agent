package fixture

import "strings"

const csvSeparator = ","

func ExportCSV(header []string, rows [][]string) string {
	lines := make([]string, 0, len(rows)+1)
	lines = append(lines, strings.Join(header, csvSeparator))
	for _, row := range rows {
		lines = append(lines, strings.Join(row, csvSeparator))
	}
	return strings.Join(lines, "\n") + "\n"
}
