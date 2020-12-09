package utils

import (
	"net"
	"strconv"
)

// GetRandomPort get a random port which is available
func GetRandomPort() (string, error) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return "", err
	}
	port := l.Addr().(*net.TCPAddr).Port

	return strconv.Itoa(port), nil
}
