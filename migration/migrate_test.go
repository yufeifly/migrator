package migration

import (
	"fmt"
	"github.com/yufeifly/migrator/model"
	"testing"
)

func TestTryMigrate(t *testing.T) {
	migrateOpts := model.MigrateOpts{
		//Container:     "9f42f4547a45", // to identify the container in source node
		ServiceID:     "",
		CheckpointID:  "cp-redis",
		CheckpointDir: "/tmp",
		Address: model.Address{
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
