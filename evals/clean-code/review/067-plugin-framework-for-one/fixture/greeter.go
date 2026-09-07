package fixture

type Greeter interface {
	Greet(name string) string
}

type greeter struct{}

func NewGreeter() Greeter {
	return greeter{}
}

func (greeter) Greet(name string) string {
	return "Hello, " + name
}
