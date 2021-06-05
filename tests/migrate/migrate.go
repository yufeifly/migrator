package main

import (
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/api/types"
	"github.com/yufeifly/migrator/client"
)

func TestMigration() {
	options := types.MigrateOpts{
		Address:       types.Address{},
		CID:           "s1.c1",
		SID:           "s1",
		CheckpointID:  "cp-redis",
		CheckpointDir: "/tmp",
	}

	cli := client.NewClient(types.Address{
		IP:   "127.0.0.1",
		Port: "6789",
	})

	err := cli.SendMigrate(options)
	if err != nil {
		logrus.Errorf("TestMigration err: %v", err)
	}
}

func main() {
	TestMigration()
}
