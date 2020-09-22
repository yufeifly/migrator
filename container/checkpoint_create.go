package container

import (
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"github.com/yufeifly/proxyd/model"
)

func CheckpointCreate(c *gin.Context) {
	if err := CreateCheckpoint(c, model.CheckpointOpts{}); err != nil {
		ReportErr(c, err)
		panic(err)
	}

	c.JSON(200, gin.H{
		"migrate": "success",
	})
}

func CreateCheckpoint(c *gin.Context, checkpointOpts model.CheckpointOpts) error {
	container := c.Request.URL.Query().Get("container")
	checkpointID := c.Request.URL.Query().Get("checkpointID")

	chOpts := types.CheckpointCreateOptions{
		CheckpointID:  checkpointID,
		CheckpointDir: checkpointOpts.CheckPointDir,
		Exit:          true,
	}

	err := cli.CheckpointCreate(ctx, container, chOpts)

	return err
}
