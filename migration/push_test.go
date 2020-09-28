package migration

import (
	"fmt"
	"github.com/yufeifly/proxyd/model"
	"testing"
)

func TestPushCheckpoint(t *testing.T) {
	MigOpts := model.MigrationOpts{
		CheckpointOpts: model.CheckpointOpts{
			CheckPointID:  "redis-cp",
			CheckPointDir: "/tmp",
		},
		DestIP:      "0.0.0.0",
		DestPort:    "6789",
		ContainerID: "85ea0420bb58",
	}
	err := PushCheckpoint(MigOpts)
	if err != nil {
		fmt.Println("err: ", err)
	}
}
