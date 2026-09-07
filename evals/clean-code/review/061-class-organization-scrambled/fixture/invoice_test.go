package fixture

import "testing"

func TestRenderGreetsAndStatesAmount(t *testing.T) {
	invoice := NewInvoice(Customer{Name: "Ada"}, 1250)

	rendered := invoice.Render()

	if rendered != "Dear Ada, you owe $12.50" {
		t.Errorf("Render = %q, want %q", rendered, "Dear Ada, you owe $12.50")
	}
}
