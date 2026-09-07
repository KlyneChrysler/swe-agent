package fixture

import (
	"net"
	"strconv"
	"strings"
)

const (
	minPort = 1
	maxPort = 65535
)

func parseHost(text string) (string, bool) {
	host := strings.TrimSpace(text)
	return host, host != ""
}

func parsePort(text string) (int, bool) {
	port, err := strconv.Atoi(text)
	if err != nil {
		return 0, false
	}
	return port, port >= minPort && port <= maxPort
}

func ListenAddress(text string) (string, bool) {
	hostText, portText, found := strings.Cut(text, ":")
	if !found {
		return "", false
	}
	host, ok := parseHost(hostText)
	if !ok {
		return "", false
	}
	port, ok := parsePort(portText)
	if !ok {
		return "", false
	}
	return net.JoinHostPort(host, strconv.Itoa(port)), true
}
