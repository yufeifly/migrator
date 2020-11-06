package container

import (
	"github.com/docker/docker/api/types"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/model"
	"github.com/yufeifly/migrator/utils"
	"os"
)

//CreateCheckpoint create a checkpoint for a container
func CreateCheckpoint(checkpointOpts model.CheckpointOpts) error {
	header := "container.CreateCheckpoint"
	chOpts := types.CheckpointCreateOptions{
		CheckpointID:  checkpointOpts.CheckPointID,
		CheckpointDir: checkpointOpts.CheckPointDir,
		Exit:          true, // todo this should be set by user
	}

	// delete the checkpoint dir if it exist
	cpPath := chOpts.CheckpointDir + "/" + chOpts.CheckpointID
	if utils.FileExist(cpPath) {
		err := os.RemoveAll(cpPath)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}

	//
	err := cli.CheckpointCreate(ctx, checkpointOpts.Container, chOpts)
	if err != nil {
		logrus.Errorf("%s, CheckpointCreate err: %v", header, err)
	}
	return err
}
