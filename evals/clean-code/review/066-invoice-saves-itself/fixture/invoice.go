package fixture

import (
	"database/sql"
	"fmt"
	"log"
)

const percent = 100

const insertInvoice = `INSERT INTO invoices (id, customer, cents) VALUES ($1, $2, $3)`

// Invoice is a bill for one customer, priced in cents before tax.
type Invoice struct {
	ID         string
	Customer   string
	Cents      int64
	TaxPercent int64
}

func (i Invoice) Total() int64 {
	return i.Cents + i.Cents*i.TaxPercent/percent
}

func (i Invoice) Save(db *sql.DB) error {
	log.Printf("saving invoice %s for %s", i.ID, i.Customer)
	if _, err := db.Exec(insertInvoice, i.ID, i.Customer, i.Cents); err != nil {
		return fmt.Errorf("save invoice %s: %w", i.ID, err)
	}
	return nil
}
