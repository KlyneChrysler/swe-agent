package fixture

type Registry struct {
	values map[string]string
}

func NewRegistry() *Registry {
	return &Registry{values: map[string]string{}}
}

func (r *Registry) Set(key, value string) bool {
	_, existed := r.values[key]
	r.values[key] = value
	return existed
}

func (r *Registry) Get(key string) (string, bool) {
	value, ok := r.values[key]
	return value, ok
}
