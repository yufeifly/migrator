package container

import (
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxyd/model"
	"github.com/yufeifly/proxyd/utils"
)

// CheckpointCreate handler for create a checkpoint for a container
func CheckpointCreate(c *gin.Context) {
	container := c.Query("container")
	checkpointID := c.Query("checkpointID")
	checkpointDIR := c.Query("checkpointDIR")

	cpOpts := model.CheckpointOpts{
		Container:     container,
		CheckPointID:  checkpointID,
		CheckPointDir: checkpointDIR,
	}
	if err := CreateCheckpoint(cpOpts); err != nil {
		utils.ReportErr(c, err)
		logrus.Panic(err)
	}

	c.JSON(200, gin.H{
		"migrate": "success",
	})
}

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
