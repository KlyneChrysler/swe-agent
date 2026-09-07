package fixture

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Warehouse struct {
	URL   string
	Token string
}

func FetchStock(warehouse Warehouse, sku string) (int, error) {
	request, err := http.NewRequest(http.MethodGet, warehouse.URL+"/stock/"+sku, http.NoBody)
	if err != nil {
		return 0, fmt.Errorf("build stock request %s: %w", sku, err)
	}
	request.Header.Set("Authorization", "Bearer "+warehouse.Token)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return 0, fmt.Errorf("fetch stock %s: %w", sku, err)
	}
	defer response.Body.Close()
	return readCount(response.Body)
}

func readCount(body io.Reader) (int, error) {
	text, err := io.ReadAll(body)
	if err != nil {
		return 0, fmt.Errorf("read stock count: %w", err)
	}
	return strconv.Atoi(strings.TrimSpace(string(text)))
}
