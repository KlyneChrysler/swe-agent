package fixture

import "net/http"

func IsUpstreamHealthy(baseURL string) bool {
	response, err := http.Head(baseURL + "/health")
	if err != nil {
		return false
	}
	response.Body.Close()
	return response.StatusCode == http.StatusOK
}
