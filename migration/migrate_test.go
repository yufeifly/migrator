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

func TestMakeNameForService(t *testing.T) {
	name := MakeNameForService("service.A1")
	fmt.Println("new name: ", name)
	if name == "service.A2" {
		fmt.Println("pass")
	} else {
		t.Error("not pass")
	}
}
