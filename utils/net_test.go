package utils

import (
	"fmt"
	"testing"
)

func TestGetRandomPort(t *testing.T) {
	port, err := GetRandomPort()
	if err != nil {
		t.Errorf("err: %v", err)
	} else {
		fmt.Println("port: ", port)
	}
}
