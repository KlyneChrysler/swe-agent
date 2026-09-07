package fixture

import (
	"fmt"
	"os"
	"path/filepath"
)

const receiptMode = 0o644

type Receipts struct {
	dir string
}

func NewReceipts(dir string) Receipts {
	return Receipts{dir: dir}
}

func (r Receipts) Store(orderID, body string) error {
	path := filepath.Join(r.dir, orderID+".txt")
	if err := os.WriteFile(path, []byte(body), receiptMode); err != nil {
		return fmt.Errorf("store receipt %s: %w", orderID, err)
	}
	return nil
}
