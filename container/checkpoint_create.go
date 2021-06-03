package container

import (
	ctypes "github.com/docker/docker/api/types"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/migrator/api/types"
	"github.com/yufeifly/migrator/utils"
	"os"
)

//CreateCheckpoint create a checkpoint for a container
func CreateCheckpoint(checkpointOpts types.CheckpointReqOpts) error {
	header := "container.CreateCheckpoint"
	chOpts := ctypes.CheckpointCreateOptions{
		CheckpointID:  checkpointOpts.CheckPointID,
		CheckpointDir: checkpointOpts.CheckPointDir,
		Exit:          true,
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
	err := dockerCli.CheckpointCreate(ctx, checkpointOpts.Container, chOpts)
	if err != nil {
		logrus.Errorf("%s, CheckpointCreate err: %v", header, err)
	}
	return err
}
