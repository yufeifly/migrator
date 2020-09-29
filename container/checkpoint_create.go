package container

import (
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"github.com/yufeifly/proxyd/model"
)

// CheckpointCreate handler for create a checkpoint for a container
func CheckpointCreate(c *gin.Context) {
	container := c.Request.URL.Query().Get("container")
	checkpointID := c.Request.URL.Query().Get("checkpointID")
	checkpointDIR := c.Request.URL.Query().Get("checkpointDIR")

	cpOpts := model.CheckpointOpts{
		Container:     container,
		CheckPointID:  checkpointID,
		CheckPointDir: checkpointDIR,
	}
	if err := CreateCheckpoint(cpOpts); err != nil {
		ReportErr(c, err)
		panic(err)
	}

	c.JSON(200, gin.H{
		"migrate": "success",
	})
}

//CreateCheckpoint create a checkpoint for a container
func CreateCheckpoint(checkpointOpts model.CheckpointOpts) error {

	chOpts := types.CheckpointCreateOptions{
		CheckpointID:  checkpointOpts.CheckPointID,
		CheckpointDir: checkpointOpts.CheckPointDir,
		Exit:          true,
	}

	err := cli.CheckpointCreate(ctx, checkpointOpts.Container, chOpts)

	return err
}
