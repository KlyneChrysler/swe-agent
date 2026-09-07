package fixture

import "testing"

func items(count int) []string {
	list := make([]string, count)
	for index := range list {
		list[index] = "item"
	}
	return list
}

func TestFirstPageIsFullWithNext(t *testing.T) {
	page := PageOf(items(25), 0)
	if len(page.Items) != 20 || !page.HasNext {
		t.Fatalf("page 0 = %d items, HasNext %v; want 20, true", len(page.Items), page.HasNext)
	}
}

func TestLastPageIsPartialWithoutNext(t *testing.T) {
	page := PageOf(items(25), 1)
	if len(page.Items) != 5 || page.HasNext {
		t.Fatalf("page 1 = %d items, HasNext %v; want 5, false", len(page.Items), page.HasNext)
	}
}

func TestPagePastEndIsEmpty(t *testing.T) {
	if got := len(PageOf(items(25), 3).Items); got != 0 {
		t.Fatalf("page 3 = %d items, want 0", got)
	}
}
