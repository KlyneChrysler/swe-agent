package fixture

import "testing"

func TestFetchCustomerFindsKnownID(t *testing.T) {
	book := CustomerBook{byID: map[string]Customer{"c1": {ID: "c1", Name: "Ada"}}}
	if customer, ok := book.FetchCustomer("c1"); !ok || customer.Name != "Ada" {
		t.Fatalf("FetchCustomer = %+v, %v; want Ada, true", customer, ok)
	}
}

func TestRetrieveOrderFindsKnownNumber(t *testing.T) {
	log := OrderLog{entries: []Order{{Number: 7, CustomerID: "c1"}}}
	if order, ok := log.RetrieveOrder(7); !ok || order.CustomerID != "c1" {
		t.Fatalf("RetrieveOrder = %+v, %v; want c1, true", order, ok)
	}
}

func TestGetProductIgnoresCase(t *testing.T) {
	catalog := Catalog{bySKU: map[string]Product{"EL-1": {SKU: "EL-1", Cents: 500}}}
	if product, ok := catalog.GetProduct("el-1"); !ok || product.Cents != 500 {
		t.Fatalf("GetProduct = %+v, %v; want 500, true", product, ok)
	}
}
