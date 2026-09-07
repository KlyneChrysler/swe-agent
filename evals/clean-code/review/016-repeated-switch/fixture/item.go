package fixture

type Kind int

const (
	Physical Kind = iota
	Digital
	Subscription
)

type Item struct {
	Kind      Kind
	BaseCents int
}
