package migration

import (
	"fmt"
	"github.com/yufeifly/migrator/api/types"
	"github.com/yufeifly/migrator/scheduler"
	"testing"
)

func TestMigrateOneWithLogging(t *testing.T) {
	//
	scheduler.Init()
	scheduler.RegisterServices()
	//
	migrateOpts := MigrateOpts{
		CID:           "s1.c1",
		SID:           "s1",
		CheckpointID:  "cp-redis",
		CheckpointDir: "/tmp",
		Address: types.Address{
			IP:   "192.168.134.135",
			Port: "6789",
		},
	}
	err := MigrateOneWithLogging(migrateOpts)
	if err != nil {
		fmt.Printf("TestTryMigrate err: %v\n", err)
	}
}
