package fixture

import (
	"fmt"
	"strings"
)

func padRight(text string, width int) string {
	if len(text) >= width {
		return text
	}
	return text + strings.Repeat(" ", width-len(text))
}

func formatCents(cents int) string {
	return fmt.Sprintf("$%d.%02d", cents/centsPerDollar, cents%centsPerDollar)
}

func formatRow(line Line) string {
	return padRight(line.Name, columnWidth) + formatCents(line.Cents)
}

type Line struct {
	Name  string
	Cents int
}

func Render(lines []Line) string {
	rows := make([]string, 0, len(lines))
	for _, line := range lines {
		rows = append(rows, formatRow(line))
	}
	return strings.Join(rows, "\n")
}

const columnWidth = 24

const centsPerDollar = 100
