package fixture

import "testing"

func TestDefaultRegistryBuildsUpperCase(t *testing.T) {
	plugin, err := DefaultRegistry().Build("upper")

	if err != nil {
		t.Fatalf("Build returned error: %v", err)
	}
	if plugin.Apply("abc") != "ABC" {
		t.Errorf("Apply = %q, want %q", plugin.Apply("abc"), "ABC")
	}
}

func TestBuildRejectsUnknownPlugin(t *testing.T) {
	_, err := DefaultRegistry().Build("lower")

	if err == nil {
		t.Fatal("Build accepted an unregistered plugin")
	}
}

func TestGreetNamesTheGuest(t *testing.T) {
	if NewGreeter().Greet("Ada") != "Hello, Ada" {
		t.Errorf("Greet = %q, want %q", NewGreeter().Greet("Ada"), "Hello, Ada")
	}
}
