package migration

import (
	"fmt"
	"github.com/yufeifly/migrator/api/types"
	"github.com/yufeifly/migrator/scheduler"
	"testing"
)

func TestMigrateOneWithLogging(t *testing.T) {

	scheduler.Init()
	scheduler.RegisterServices()

	migrateOpts := types.MigrateOpts{
		CID:           "s1.c1",
		SID:           "s1",
		CheckpointID:  "cp-redis",
		CheckpointDir: "/tmp",
		Address: types.Address{
			IP:   "192.168.134.135", // target address
			Port: "6789",
		},
	}

	err := MigrateOneWithLogging(migrateOpts)
	if err != nil {
		fmt.Printf("TestMigrateOneWithLogging err: %v\n", err)
	}
}
