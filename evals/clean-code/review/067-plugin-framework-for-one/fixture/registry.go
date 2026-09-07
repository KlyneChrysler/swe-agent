package fixture

import (
	"fmt"
	"strings"
)

// Plugin transforms a line of text before it is written to the report.
type Plugin interface {
	Apply(text string) string
}

type Registry struct {
	factories map[string]func() Plugin
}

func NewRegistry() *Registry {
	return &Registry{factories: map[string]func() Plugin{}}
}

func (r *Registry) Register(name string, factory func() Plugin) {
	r.factories[name] = factory
}

func (r *Registry) Build(name string) (Plugin, error) {
	factory, ok := r.factories[name]
	if !ok {
		return nil, fmt.Errorf("build plugin %q: not registered", name)
	}
	return factory(), nil
}

func DefaultRegistry() *Registry {
	registry := NewRegistry()
	registry.Register("upper", func() Plugin { return upperCase{} })
	return registry
}

type upperCase struct{}

func (upperCase) Apply(text string) string {
	return strings.ToUpper(text)
}
