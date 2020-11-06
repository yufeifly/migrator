package client

import (
	"fmt"
	"testing"
)

func TestClient_ConsumeAdder(t *testing.T) {
	cli := NewClient()
	err := cli.ConsumedAdder()
	if err != nil {
		t.Errorf("cli.ConsumeAdder failed, err: %v", err)
	} else {
		fmt.Println("pass")
	}
}
