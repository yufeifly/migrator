package utils

import (
	"fmt"
	"testing"
)

func TestMakeNameForService(t *testing.T) {
	name := MakeNameForService("service.A1")
	fmt.Println("new name: ", name)
	if name == "service.A2" {
		fmt.Println("pass")
	} else {
		t.Error("not pass")
	}
}
