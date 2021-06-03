package migration

import (
	"fmt"
	"github.com/yufeifly/migrator/api/types"
	"testing"
)

func TestTryMigrate(t *testing.T) {
	migrateOpts := MigrateOpts{
		CID:           "9f42f4547a45", // to identify the container in source node
		SID:           "",
		CheckpointID:  "cp-redis",
		CheckpointDir: "/tmp",
		Address: types.Address{
			IP:   "127.0.0.1",
			Port: "6789",
		},
	}
	err := TryMigrate(migrateOpts)
	if err != nil {
		fmt.Printf("TestTryMigrate err: %v\n", err)
	} else {
		fmt.Printf("TestTryMigrate pass\n")
	}
}
