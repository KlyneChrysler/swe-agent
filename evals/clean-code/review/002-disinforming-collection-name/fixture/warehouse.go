package fixture

type Warehouse struct {
	skuList map[string]int
}

func NewWarehouse() *Warehouse {
	return &Warehouse{skuList: map[string]int{}}
}

func (w *Warehouse) Receive(sku string, quantity int) {
	w.skuList[sku] += quantity
}

func (w *Warehouse) OnHand(sku string) int {
	return w.skuList[sku]
}
