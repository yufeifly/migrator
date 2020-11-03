package utils

import (
	"net"
	"strconv"
)

// get a random port which is available
func GetRandomPort() (string, error) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return "", err
	}
	port := l.Addr().(*net.TCPAddr).Port

	return strconv.Itoa(port), nil

}
