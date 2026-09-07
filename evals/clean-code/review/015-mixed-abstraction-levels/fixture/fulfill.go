package fixture

import "strings"

type Order struct {
	ID           string
	Recipient    string
	AddressLines []string
	SKUs         []string
}

type Stock interface {
	Reserve(skus []string) error
}

type Courier interface {
	Dispatch(orderID, label string) error
}

type Fulfiller struct {
	stock   Stock
	courier Courier
}

func NewFulfiller(stock Stock, courier Courier) Fulfiller {
	return Fulfiller{stock: stock, courier: courier}
}

func (f Fulfiller) Fulfill(order Order) error {
	if err := f.stock.Reserve(order.SKUs); err != nil {
		return err
	}
	label := strings.ToUpper(strings.TrimSpace(order.Recipient)) + "\n" + strings.Join(order.AddressLines, "\n")
	return f.courier.Dispatch(order.ID, label)
}
