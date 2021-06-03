package migration

import (
	"fmt"
	"github.com/yufeifly/migrator/api/types"
	"testing"
)

func TestPushCheckpoint(t *testing.T) {
	PushOpts := PushOpts{
		CheckPointID:  "redis-cp",
		CheckPointDir: "/tmp",
		Dest:          types.Address{},
		CID:           "",
		SID:           "",
		Port:          "",
	}
	err := PushCheckpoint(PushOpts)
	if err != nil {
		fmt.Println("err: ", err)
	}
}
