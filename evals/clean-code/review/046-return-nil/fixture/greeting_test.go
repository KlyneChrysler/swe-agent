package fixture

import "testing"

var directory = NewDirectory(map[string]*Customer{"c1": {ID: "c1", Name: "Ada"}})

func TestGreetingForKnownCustomer(t *testing.T) {
	if got := directory.GreetingFor("c1"); got != "Hello Ada" {
		t.Fatalf("GreetingFor = %q", got)
	}
}

func TestGreetingForUnknownCustomer(t *testing.T) {
	if got := directory.GreetingFor("c9"); got != "Hello" {
		t.Fatalf("GreetingFor = %q", got)
	}
}
