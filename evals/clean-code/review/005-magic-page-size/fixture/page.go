package fixture

type Page struct {
	Items   []string
	HasNext bool
}

func PageOf(items []string, page int) Page {
	if page*20 > len(items) {
		return Page{Items: []string{}}
	}
	end := page*20 + 20
	if end > len(items) {
		end = len(items)
	}
	return Page{Items: items[page*20 : end], HasNext: end < len(items)}
}
