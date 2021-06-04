package client

import (
	"github.com/yufeifly/migrator/api/types"
	"testing"
)

func TestClient_ConsumeAdder(t *testing.T) {
	cli := NewClient(types.Address{
		IP:   "127.0.0.1",
		Port: "6789",
	})
	err := cli.ConsumedAdder("s1.c1")
	if err != nil {
		t.Errorf("cli.ConsumeAdder failed, err: %v", err)
	}
}
