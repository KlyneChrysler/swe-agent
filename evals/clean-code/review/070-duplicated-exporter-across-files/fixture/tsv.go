package fixture

import "strings"

const tsvSeparator = "\t"

func ExportTSV(header []string, rows [][]string) string {
	lines := make([]string, 0, len(rows)+1)
	lines = append(lines, strings.Join(header, tsvSeparator))
	for _, row := range rows {
		lines = append(lines, strings.Join(row, tsvSeparator))
	}
	return strings.Join(lines, "\n") + "\n"
}
