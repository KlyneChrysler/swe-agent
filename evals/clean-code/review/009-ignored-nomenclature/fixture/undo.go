package fixture

type Pile struct {
	items []string
}

func (p *Pile) Put(item string) {
	p.items = append(p.items, item)
}

func (p *Pile) Take() (string, bool) {
	if len(p.items) == 0 {
		return "", false
	}
	last := len(p.items) - 1
	item := p.items[last]
	p.items = p.items[:last]
	return item, true
}

func (p *Pile) Size() int {
	return len(p.items)
}
