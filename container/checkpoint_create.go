package container

import (
	"github.com/docker/docker/api/types"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/model"
)

//CreateCheckpoint create a checkpoint for a container
func CreateCheckpoint(checkpointOpts model.CheckpointOpts) error {
	header := "container.CreateCheckpoint"
	chOpts := types.CheckpointCreateOptions{
		CheckpointID:  checkpointOpts.CheckPointID,
		CheckpointDir: checkpointOpts.CheckPointDir,
		Exit:          true, // todo this should be set by user
	}

	err := cli.CheckpointCreate(ctx, checkpointOpts.Container, chOpts)
	if err != nil {
		logrus.Errorf("%s, CheckpointCreate err: %v", header, err)
	}
	return err
}
