package container

import (
	"context"
	"fmt"
	"github.com/yufeifly/proxyd/model"

	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
)

func Start(c *gin.Context) {
	containerID := c.Request.URL.Query().Get("containerId")
	checkpointID := c.Request.URL.Query().Get("checkpointID")
	checkpointDir := c.Request.URL.Query().Get("checkpointDir")

	startOpts := model.StartOpts{
		CStartOpts: types.ContainerStartOptions{
			CheckpointID:  checkpointID,
			CheckpointDir: checkpointDir,
		},
		ContainerID: containerID,
	}

	err := startContainer(startOpts)
	if err != nil {

	}
	c.JSON(200, gin.H{
		"result": "success",
	})
}

func startContainer(startOpts model.StartOpts) error {
	err := cli.ContainerStart(context.Background(), startOpts.ContainerID, startOpts.CStartOpts)
	if err != nil {
		fmt.Println("start container failed")
		return err
	}
	return nil
}
