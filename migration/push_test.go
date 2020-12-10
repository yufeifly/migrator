package migration

import (
	"fmt"
	"testing"
)

func TestPushCheckpoint(t *testing.T) {
	PushOpts := PushOpts{
		CheckPointID:  "redis-cp",
		CheckPointDir: "/tmp",
		DestIP:        "0.0.0.0",
		DestPort:      "6789",
		ContainerID:   "85ea0420bb58",
	}
	err := PushCheckpoint(PushOpts)
	if err != nil {
		fmt.Println("err: ", err)
	}
}
